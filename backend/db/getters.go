package db

import (
	"context"
	"fmt"
	"log"

	"github.com/CanobbioE/reelo/backend/dto"
)

// PlayerID retrieves a player id from the database given its name and surname
func (database *DB) PlayerID(ctx context.Context, name, surname string) (int, error) {
	var id int
	q := findPlayerIDByNameAndSurname
	err := database.db.QueryRow(q, name, surname).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("Error getting player id: %v", err)
	}
	return id, nil
}

// Results retrives all the results a player had in all the years he partecipated
func (database *DB) Results(ctx context.Context, name, surname string) ([]Result, error) {
	var results []Result
	q := findResultsByNameAndSurname
	rows, err := database.db.QueryContext(ctx, q, name, surname)
	if err != nil {
		return results, fmt.Errorf("Error getting results: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Result{}
		err := rows.Scan(&r.Time, &r.Exercises, &r.Score, &r.Year, &r.Category)
		if err != nil {
			return results, fmt.Errorf("Error getting results: %v", err)
		}
		results = append(results, r)
	}
	return results, nil
}

// AllPlayers retrieves all players from the database
func (database *DB) AllPlayers(ctx context.Context) ([]Player, error) {
	var players []Player
	q := findAllPlayers
	rows, err := database.db.QueryContext(ctx, q)
	if err != nil {
		return players, fmt.Errorf("Error getting players: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		err := rows.Scan(&p.Name, &p.Surname)
		if err != nil {
			return players, fmt.Errorf("Error getting players: %v", err)
		}
		players = append(players, p)
	}
	return players, nil
}

// PlayerPartecipationYears retrieves a list of all the years a player has played
func (database *DB) PlayerPartecipationYears(ctx context.Context, name, surname string) ([]int, error) {
	var years []int
	pID, err := database.PlayerID(ctx, name, surname)
	if err != nil {
		return years, err
	}
	q := findPartecipationYearsByPlayer
	rows, err := database.db.QueryContext(ctx, q, pID)
	if err != nil {
		return years, fmt.Errorf("Error getting partcipation years: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var y int
		err := rows.Scan(&y)
		if err != nil {
			return years, fmt.Errorf("Error getting partcipation years: %v", err)
		}
		years = append(years, y)
	}
	return years, nil
}

// Score retrieve the score of a given player for a given year
func (database *DB) Score(name, surname string, year int, isParis bool) (float64, error) {
	var score float64
	q := findScoresByPlayerAndYear

	q = adaptToParis(q, isParis)

	err := database.db.QueryRow(q, name, surname, year).Scan(&score)
	if err != nil {
		return score, fmt.Errorf("Error getting scores: %v", (err))
	}
	return score, nil
}

// Exercises retrieve the number of exercises the specified player
// has completed in a given year
func (database *DB) Exercises(name, surname string, year int, isParis bool) (int, error) {
	var es int
	q := findExercisesByPlayerAndYear

	adaptToParis(q, isParis)

	err := database.db.QueryRow(q, name, surname, year).Scan(&es)
	if err != nil {
		return es, fmt.Errorf("Error getting exercises: %v", (err))
	}
	return es, nil
}

// Categories returns the categories in which the given player
// has partecipated in the specified year.
// There could be more than one category for a year,
// this could happen for namesakes or when we have international results
func (database *DB) Categories(ctx context.Context, name, surname string, year int) ([]string, error) {
	var categories []string
	q := findCategoriesByPlayerAndYear
	rows, err := database.db.QueryContext(ctx, q, name, surname, year)
	if err != nil {
		return categories, fmt.Errorf("Error getting categories: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			return categories, fmt.Errorf("Error getting categories: %v", err)
		}
		categories = append(categories, c)
	}
	return categories, nil

}

// AvgScoresOfCategories calculates the average scores for
// all the categories in the given year
func (database *DB) AvgScoresOfCategories(year int) (float64, error) {
	var avg float64
	q := findAvgScoresByYear
	err := database.db.QueryRow(q, year).Scan(&avg)
	if err != nil {
		return avg, fmt.Errorf("Error getting average score of categories: %v", err)
	}
	return avg, nil
}

// AvgScore returns the score's average for the given year and category
func (database *DB) AvgScore(year int, category string) (float64, error) {
	var avg float64
	q := findAvgScoresByYearAndCategory
	err := database.db.QueryRow(q, year, category).Scan(&avg)
	if err != nil {
		return avg, fmt.Errorf("Error getting average score: %v", err)
	}

	return avg, nil
}

// MaxScore calculates the maximum score obtained by any player
// in the given year and category
func (database *DB) MaxScore(year int, category string) (float64, error) {
	var max float64
	q := findMaxScoreByYearAndCategory
	err := database.db.QueryRow(q, year, category).Scan(&max)
	if err != nil {
		return max, fmt.Errorf("Error getting max score: %v", err)
	}
	return max, nil
}

// Password retrieve the hashed password of a user
// password must be saved by going through:
//
// hashPassword := sha256.New()
// hashPassword.Write([]byte("password"))
// return string(hashPassword.Sum(nil))
func (database DB) Password(ctx context.Context, username string) (string, error) {
	var hash string
	q := findPasswordByUsername
	err := database.db.QueryRow(q, username).Scan(&hash)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

// ReeloCostants retrieve the costants used to calculate the reelo score
func (database *DB) ReeloCostants() (Costants, error) {
	c := Costants{}
	q := findAllCostants

	err := database.db.QueryRow(q).Scan(&c.StartingYear, &c.ExercisesCostant, &c.PFinal, &c.MultiplicativeFactor, &c.AntiExploit, &c.NoPartecipationPenalty)
	if err != nil {
		return c, err
	}
	return c, nil
}

// LastKnownYear returns the last year we know anything about
func (database *DB) LastKnownYear() (int, error) {
	var year int
	q := findMaxYear
	err := database.db.QueryRow(q).Scan(&year)
	if err != nil {
		return year, fmt.Errorf("Error getting last known year: %v", err)
	}
	return year, nil
}

// LastKnownCategoryForPlayer returns the last category into which
// we know the specified player has partecipated
func (database *DB) LastKnownCategoryForPlayer(name, surname string) (string, error) {
	var category string
	q := findCategoryByPlayerAndYear
	err := database.db.QueryRow(q, name, surname, name, surname).Scan(&category)
	if err != nil {
		return category, fmt.Errorf("Error getting last known category: %v", err)
	}
	return category, nil
}

// IsResultFromParis checks if the result associated w/ the given data
// comes from an international game
func (database *DB) IsResultFromParis(ctx context.Context, name, surname string,
	year int, category string) (bool, error) {

	q := findCityByPlayerAndYearAndCategory
	rows, err := database.db.QueryContext(ctx, q, name, surname, year, category)
	if err != nil {
		return false, fmt.Errorf("Error getting results from paris: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			return false, fmt.Errorf("Error getting results from paris: %v", err)
		}
		if c == "paris" {
			return true, nil
		}
	}
	return false, nil
}

// AllRanks returns a list of all the player and ranks inside the database.
// A rank is composed by a player's name, surname, reelo and last category
// into which he has played
func (database *DB) AllRanks(ctx context.Context) (ranks []dto.Rank, err error) {
	q := findAllPlayersRanks
	rows, err := database.db.QueryContext(ctx, q)
	if err != nil {
		log.Printf("Error getting all ranks: %v", err)
		return ranks, err
	}
	defer rows.Close()

	for rows.Next() {
		var r dto.Rank
		err := rows.Scan(&r.Name, &r.Surname, &r.Category, &r.Reelo)
		if err != nil {
			log.Printf("Error getting all ranks: %v", err)
			return ranks, err
		}
		ranks = append(ranks, r)
	}
	return ranks, nil

}
