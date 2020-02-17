package interactor

import (
	"context"
	"io"
	"strings"

	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils/parse"
)

// Logger is the interface for the logging utility
type Logger interface {
	Log(msg string, args ...interface{})
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
}

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
func (i *Interactor) ParseFileWithInfo(fileReader io.Reader, game domain.Game, format string) error {

	var results []parse.LineInfo

	category := strings.ToUpper(game.Category)
	parsedFmt, err := parse.NewFormat(strings.Split(format, " "))
	if err != nil {
		return err
	}

	// this is the warning we want to return to the front end
	results, warning := parse.File(fileReader, parsedFmt, game.Year, category)
	if warning != nil {
		i.Logger.Log("parse.File() returned warning: %v\n", warning)
		return warning
	}

	i.Logger.Log("File parsed succesfully\n\n")
	i.GameRepository.InserRankingFile(context.Background(), results, game)
	i.Logger.Log("File inserted succesfully\n\n")
	return nil
}

// DeleteIfAlreadyExists search for results from the year+category contained in
// info. If the year+category exists in the database, all the results gets erased.
func (i *Interactor) DeleteIfAlreadyExists(game domain.Game) error {

	id, err := i.GameRepository.FindIDByYearAndCategoryAndIsParis(context.Background(), game.Year, game.Category, game.IsParis)
	if err != nil {
		return err
	}

	if id != -1 {
		if err := i.ResultRepository.DeleteByGameID(context.Background(), id); err != nil {
			return err
		}
		i.Logger.Log("Deleted results with id: %v\n", id)
	} else {
		i.Logger.Log("Nothing to delete")
	}
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

// InsertRankingFile inserts all the result contained in the already parsed file into the database by making the correct calls
func (i *Interactor) InsertRankingFile(ctx context.Context, file []parse.LineInfo, game domain.Game) error {
	/*
		gamesID, err := database.Add(ctx, "giochi",
			gameInfo.Year, gameInfo.Category, gameInfo.Start, gameInfo.End, isParis)
		if err != nil {
			return err
		}
		for _, line := range file {
			if line.Name == "" && line.Surname == "" {
				continue
			}
			city := line.City
			if isParis {
				city = "paris"
			}
			var playerID int
			if !database.ContainsPlayer(ctx, line.Name, line.Surname) {
				accent := CreateAccent(gameInfo.Year, 0, city)
				playerID, err = database.Add(ctx, "giocatore", line.Name, line.Surname, accent)
				if err != nil {
					return err
				}
			}
			playerID, err = database.PlayerID(ctx, line.Name, line.Surname)
			if err != nil {
				return err
			}
			resultsID, err := database.Add(ctx, "risultato", line.Time, line.Exercises, line.Points, line.Position, 0)
			if err != nil {
				return err
			}
			_, err = database.Add(ctx, "partecipazione", playerID, gamesID, resultsID, city)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}