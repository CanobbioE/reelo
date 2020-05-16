package interactor

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/pkg/parse"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Logger is the interface for the logging utility,
// having the interface definition here allows for dependency injection
// and proper logging even on low level code.
type Logger interface {
	Log(msg string, args ...interface{})
}

// Interactor is used to interact with the externally injected repositories
type Interactor struct {
	CommentRepository       domain.CommentRepository
	CostantsRepository      domain.CostantsRepository
	GameRepository          domain.GameRepository
	ParticipationRepository domain.ParticipationRepository
	PlayerRepository        domain.PlayerRepository
	ResultRepository        domain.ResultRepository
	UserRepository          usecases.UserRepository
	HistoryRepository       usecases.HistoryRepository
	Logger                  Logger
}

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
// TODO: store the entities in an in-memory db and fix the data, then store into the actual db
func (i *Interactor) ParseFileWithInfo(fileReader io.Reader, game domain.Game, format, city string) utils.Error {
	ctx := context.Background()
	var lines []parse.LineInfo

	category := strings.ToUpper(game.Category)
	parsedFmt, err := parse.NewFormat(strings.Split(format, " "))
	if err != nil {
		return utils.NewError(err, "E_PARSE_FMT", 400)
	}

	// this is the warning we want to return to the front end
	lines, warning := parse.File(fileReader, parsedFmt, game.Year, category)
	if warning != nil {
		i.Logger.Log("ParseFileWithInfo: parse.File warning: %v", warning)
		return utils.NewError(warning, "E_PARSE_WARN", 500)
	}
	i.Logger.Log("File parsed successfully\n")

	gamesID, err := i.GameRepository.Store(ctx, game)
	if err != nil {
		i.Logger.Log("ParseFileWithInfo: cannot store game: %v", err)
		return utils.NewError(err, "E_DB_STORE", 500)
	}
	for _, line := range lines {
		if line.Name == "" && line.Surname == "" {
			continue
		}

		var playerID int
		if !i.PlayerRepository.CheckExistenceByNameAndSurname(ctx, line.Name, line.Surname) {
			accent := fmt.Sprintf("%d %s %d", line.Year, line.City, 0)
			p := domain.Player{
				Name:    line.Name,
				Surname: line.Surname,
				Accent:  accent,
				Reelo:   0,
			}
			playerID64, err := i.PlayerRepository.Store(ctx, p)
			if err != nil {
				i.Logger.Log("ParseFileWithInfo: cannot store player %v, %v: %v", p.Name, p.Surname, warning)
				return utils.NewError(err, "E_DB_STORE", 500)
			}
			playerID = int(playerID64)
		}

		playerID, err = i.PlayerRepository.FindIDByNameAndSurname(ctx, line.Name, line.Surname)
		if err != nil {
			i.Logger.Log("ParseFileWithInfo: cannot find player id: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}

		r := domain.Result{
			Time:        line.Time,
			Exercises:   line.Exercises,
			Score:       line.Points,
			Position:    line.Position,
			PseudoReelo: 0,
		}

		resultsID, err := i.ResultRepository.Store(ctx, r)
		if err != nil {
			i.Logger.Log("ParseFileWithInfo: cannot store result: %v", err)
			return utils.NewError(err, "E_DB_STORE", 500)
		}

		p := domain.Participation{
			Player: domain.Player{ID: playerID},
			Game:   domain.Game{ID: int(gamesID)},
			Result: domain.Result{ID: int(resultsID)},
			City:   line.City,
		}

		if _, err := i.ParticipationRepository.Store(ctx, p); err != nil {
			i.Logger.Log("ParseFileWithInfo: cannot store participation: %v", err)
			return utils.NewError(err, "E_DB_STORE", 500)
		}
	}

	i.Logger.Log("File inserted successfully\n")
	return utils.NewNilError()
}

// DeleteIfAlreadyExists search for results from the year+category contained in
// info. If the year+category exists in the database, all the results gets erased.
func (i *Interactor) DeleteIfAlreadyExists(game domain.Game) utils.Error {

	id, err := i.GameRepository.FindIDByYearAndCategoryAndIsParis(context.Background(), game.Year, game.Category, game.IsParis)
	if err != nil {
		if err.Error() != "no values in result set" {
			i.Logger.Log("DeleteIfAlreadyExists: cannot find game's ID %d: %v", err)
			return utils.NewError(err, "E_DB_FIND", 500)
		}
		i.Logger.Log("Nothing to delete")
		return utils.NewNilError()
	}

	if err := i.ResultRepository.DeleteByGameID(context.Background(), id); err != nil {
		i.Logger.Log("DeleteIfAlreadyExists: cannot delete result: %v", err)
		return utils.NewError(err, "E_DB_DELETE", 500)
	}
	i.Logger.Log("Deleted results with id: %v\n", id)
	return utils.NewNilError()
}

// DoesRankExist is called to verify if a year-category ranking file has been already uploaded
func (i *Interactor) DoesRankExist(year int, category string, isParis bool) (bool, utils.Error) {

	id, err := i.GameRepository.FindIDByYearAndCategoryAndIsParis(context.Background(), year, category, isParis)
	if err != nil {
		if strings.Contains(err.Error(), "no values in result set") {
			return false, utils.NewNilError()
		}
		i.Logger.Log("DoesRankExist: cannot find game's ID: %v", err)
		return false, utils.NewError(err, "E_DB_FIND", 500)
	}

	return id != -1, utils.NewNilError()
}

// Log allows for logging on low level code
func (i *Interactor) Log(msg string, args ...interface{}) {
	i.Logger.Log(msg, args...)
}
