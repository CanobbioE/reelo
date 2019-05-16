package db

import (
	"context"
	"log"
)

// PlayerID retrieves a player id from the database given its name and surname
func (database *DB) PlayerID(ctx context.Context, name, surname string) (id int) {
	q := findPlayerIDByNameAndSurname
	err := database.db.QueryRow(q, name, surname).Scan(&id)
	if err != nil {
		log.Printf("Error getting player id: %v", err)
		return id
	}
	return id
}

// Results retrives all the results a player had in all the years he partecipated
func (database *DB) Results(ctx context.Context, name, surname string) (results []Result) {
	q := findResultsByNameAndSurname
	rows, err := database.db.QueryContext(ctx, q, name, surname)
	if err != nil {
		log.Printf("Error getting results: %v", err)
		return results
	}
	defer rows.Close()

	for rows.Next() {
		r := Result{}
		err := rows.Scan(&r.Time, &r.Exercises, &r.Score, &r.Year, &r.Category)
		if err != nil {
			log.Printf("Error getting results: %v", err)
			return results
		}
		results = append(results, r)
	}
	return results
}

// AllPlayers retrieves all players from the database
func (database *DB) AllPlayers(ctx context.Context) (players []Player) {
	q := findAllPlayers
	rows, err := database.db.QueryContext(ctx, q)
	if err != nil {
		log.Printf("Error getting players: %v", err)
		return players
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		err := rows.Scan(&p.Name, &p.Surname)
		if err != nil {
			log.Printf("Error getting players: %v", err)
			return players
		}
		players = append(players, p)
	}
	return players
}

// PlayerPartecipationYears retrieves a list of all the years a player has played
func (database *DB) PlayerPartecipationYears(ctx context.Context, name, surname string) (years []int) {
	pID := database.PlayerID(ctx, name, surname)
	q := findPartecipationYearsByPlayer
	rows, err := database.db.QueryContext(ctx, q, pID)
	if err != nil {
		log.Printf("Error getting partcipation years: %v", err)
		return years
	}
	defer rows.Close()

	for rows.Next() {
		var y int
		err := rows.Scan(&y)
		if err != nil {
			log.Printf("Error getting partcipation years: %v", err)
			return years
		}
		years = append(years, y)
	}
	return years
}

// Score retrieve the score of a given player for a given year
func (database *DB) Score(name, surname string, year int, isParis bool) (score float64) {
	q := findScoresByPlayerAndYear

	q = adaptToParis(q, isParis)

	err := database.db.QueryRow(q, name, surname, year).Scan(&score)
	if err != nil {
		log.Printf("Error getting scores: %v", (err))
		return score
	}
	return score
}

// Exercises retrieve the number of exercises the specified player
// has completed in a given year
func (database *DB) Exercises(name, surname string, year int, isParis bool) (es int) {
	q := findExercisesByPlayerAndYear

	adaptToParis(q, isParis)

	err := database.db.QueryRow(q, name, surname, year).Scan(&es)
	if err != nil {
		log.Printf("Error getting exercises: %v", (err))
		return es
	}
	return es
}

// Categories returns the categories in which the given player
// has partecipated in the specified year.
// There could be more than one category for a year,
// this could happen for namesakes or when we have international results
func (database *DB) Categories(ctx context.Context, name, surname string, year int) (categories []string) {
	q := findCategoriesByPlayerAndYear
	rows, err := database.db.QueryContext(ctx, q, name, surname, year)
	if err != nil {
		log.Printf("Error getting categories: %v", err)
		return categories
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			log.Printf("Error getting categories: %v", err)
			return categories
		}
		categories = append(categories, c)
	}
	return categories

}

// AvgScoresOfCategories calculates the average scores for
// all the categories in the given year
func (database *DB) AvgScoresOfCategories(year int) (avg float64) {
	q := findAvgScoresByYear
	err := database.db.QueryRow(q, year).Scan(&avg)
	if err != nil {
		log.Printf("Error getting average score of categories: %v", err)
		return avg
	}
	return avg
}

// AvgScore returns the score's average for the given year and category
func (database *DB) AvgScore(year int, category string) (avg float64) {
	q := findAvgScoresByYearAndCategory
	err := database.db.QueryRow(q, year, category).Scan(&avg)
	if err != nil {
		log.Printf("Error getting average score: %v", err)
		return avg
	}

	return avg
}

// MaxScore calculates the maximum score obtained by any player
// in the given year and category
func (database *DB) MaxScore(year int, category string) (max float64) {
	q := findMaxScoreByYearAndCategory
	err := database.db.QueryRow(q, year, category).Scan(&max)
	if err != nil {
		log.Printf("Error getting max score: %v", err)
		return max
	}
	return max
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
func (database *DB) LastKnownYear() (year int) {
	q := findMaxYear
	err := database.db.QueryRow(q).Scan(&year)
	if err != nil {
		log.Printf("Error getting last known year: %v", err)
		return year
	}
	return year
}

// LastKnownCategoryForPlayer returns the last category into which
// we know the specified player has partecipated
func (database *DB) LastKnownCategoryForPlayer(name, surname string) (category string) {
	q := findCategoryByPlayerAndYear
	err := database.db.QueryRow(q, name, surname, name, surname).Scan(&category)
	if err != nil {
		log.Printf("Error getting last known category: %v", err)
		return category
	}
	return category
}

// IsResultFromParis checks if the result associated w/ the given data
// comes from an international game
func (database *DB) IsResultFromParis(ctx context.Context, name, surname string,
	year int, category string) bool {

	q := findCityByPlayerAndYearAndCategory
	rows, err := database.db.QueryContext(ctx, q, name, surname, year, category)
	if err != nil {
		log.Printf("Error getting results from paris: %v", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			log.Printf("Error getting results from paris: %v", err)
			return false
		}
		if c == "paris" {
			return true
		}
	}
	return false
}
