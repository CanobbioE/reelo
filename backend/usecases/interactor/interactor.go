package interactor

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// Logger is the interface for the logging utility,
// having the interface definition here allows for dependency injection
// and proper logging even on low level code.
type Logger interface {
	Log(msg string, args ...interface{})
}

// ErrorHandler is the interface for the error handling utility,
// having the interface definition here allows for dependency injection
// and proper error handling even on low level code.
type ErrorHandler interface {
	NewError(err error, code string, httpStatus int) Error
}

// Error is the interface for the application custom error
type Error interface {
	Code() string
	Status() int
	Message() string
	String() string
}

// Interactor is used to interact with the externally injected repositories
type Interactor struct {
	CommentRepository       domain.CommentRepository
	CostantsRepository      domain.CostantsRepository
	GameRepository          domain.GameRepository
	PartecipationRepository domain.PartecipationRepository
	PlayerRepository        domain.PlayerRepository
	ResultRepository        domain.ResultRepository
	UserRepository          usecases.UserRepository
	HistoryRepository       usecases.HistoryRepository
	Logger                  Logger
	ErrorHandler            ErrorHandler
}

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
func (i *Interactor) ParseFileWithInfo(fileReader io.Reader, game domain.Game, format, city string) error {
	ctx := context.Background()
	var lines []parse.LineInfo

	category := strings.ToUpper(game.Category)
	parsedFmt, err := parse.NewFormat(strings.Split(format, " "))
	if err != nil {
		return err
	}

	// this is the warning we want to return to the front end
	lines, warning := parse.File(fileReader, parsedFmt, game.Year, category)
	if warning != nil {
		i.Logger.Log("parse.File() returned warning: %v\n", warning)
		return warning
	}
	i.Logger.Log("File parsed succesfully\n")

	gamesID, err := i.GameRepository.Store(ctx, game)
	if err != nil {
		return err
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
				return err
			}
			playerID = int(playerID64)
		}

		playerID, err = i.PlayerRepository.FindIDByNameAndSurname(ctx, line.Name, line.Surname)
		if err != nil {
			return err
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
			return err
		}

		p := domain.Partecipation{
			Player: domain.Player{ID: playerID},
			Game:   domain.Game{ID: int(gamesID)},
			Result: domain.Result{ID: int(resultsID)},
			City:   line.City,
		}

		if _, err := i.PartecipationRepository.Store(ctx, p); err != nil {
			return err
		}
	}

	i.Logger.Log("File inserted succesfully\n")
	return nil
}

// DeleteIfAlreadyExists search for results from the year+category contained in
// info. If the year+category exists in the database, all the results gets erased.
func (i *Interactor) DeleteIfAlreadyExists(game domain.Game) error {

	id, err := i.GameRepository.FindIDByYearAndCategoryAndIsParis(context.Background(), game.Year, game.Category, game.IsParis)
	if err != nil {
		if err.Error() != "no values in result set" {
			return err
		}
		i.Logger.Log("Nothing to delete")
		return nil
	}

	if err := i.ResultRepository.DeleteByGameID(context.Background(), id); err != nil {
		return err
	}
	i.Logger.Log("Deleted results with id: %v\n", id)
	return nil
}

// DoesRankExist is called to verify if a year-category ranking file has been already uploaded
func (i *Interactor) DoesRankExist(year int, category string, isParis bool) (bool, error) {

	id, err := i.GameRepository.FindIDByYearAndCategoryAndIsParis(context.Background(), year, category, isParis)
	if err != nil {
		return false, err
	}

	return id != -1, nil
}

// Log allows for logging on low level code
func (i *Interactor) Log(msg string, args ...interface{}) {
	i.Logger.Log(msg, args...)
}

// Error allows for error handling on low level code
// Technically we the interactor shouldn't be aware of the httpStatus
// because the interactor shouldn't be limited to http requests.
func (i *Interactor) Error(err error, code string, httpStatus int) string {
	return i.ErrorHandler.NewError(err, code, httpStatus).String()
}
