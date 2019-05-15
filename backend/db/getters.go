package db

import (
	"context"
	"log"
)

// GetPlayerID retrieves a player id from the database given its name and surname
func (database *DB) GetPlayerID(ctx context.Context, name, surname string) (id int) {
	q := `SELECT id FROM Giocatore WHERE nome = ? AND cognome = ?`
	err := database.db.QueryRow(q, name, surname).Scan(&id)
	if err != nil {
		log.Printf("Error getting player id: %v", err)
		return id
	}
	return id
}

// GetResults retrives all the results a player had in all the years he partecipated
func (database *DB) GetResults(ctx context.Context, name, surname string) (results []Result) {
	q := `
SELECT R.tempo, R.esercizi, R.punteggio, G.anno, G.categoria
FROM Giocatore U
JOIN Partecipazione P ON P.giocatore = U.id
JOIN Risultato R ON R.id = P.risultato
JOIN Giochi G ON G.id = P.giochi
WHERE U.Nome = ? AND U.Cognome = ?
`
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

// GetAllPlayers retrieves all players from the database
func (database *DB) GetAllPlayers(ctx context.Context) (players []Player) {
	q := `SELECT nome, cognome FROM Giocatore`
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

// GetPlayerPartecipationYears retrieves a list of all the years a player has played
func (database *DB) GetPlayerPartecipationYears(ctx context.Context, name, surname string) (years []int) {
	pID := database.GetPlayerID(ctx, name, surname)
	q := `
SELECT DISTINCT G.anno FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.id = ?
`
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

// GetScore retrieve the score of a given player for a given year
func (database *DB) GetScore(name, surname string, year int) (score float64) {
	q := `
SELECT R.punteggio FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.Cognome = ? AND G.anno = ?
`
	err := database.db.QueryRow(q, name, surname, year).Scan(&score)
	if err != nil {
		log.Printf("Error getting scores: %v", (err))
		return score
	}
	return score
}

// GetExercises retrieve the number of exercises the specified player
// has completed in a given year
func (database *DB) GetExercises(name, surname string, year int) (es int) {
	q := `
SELECT R.esercizi FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.Cognome = ? AND G.anno = ?
`
	err := database.db.QueryRow(q, name, surname, year).Scan(&es)
	if err != nil {
		log.Printf("Error getting exercises: %v", (err))
		return es
	}
	return es
}

// GetCategory returns the category in which the given player
// has partecipated in the specified year
func (database *DB) GetCategory(name, surname string, year int) (category string) {
	q := `
SELECT G.categoria FROM Giochi
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE G.anno = ? AND U.nome = ? AND U.cognome = ?
`
	err := database.db.QueryRow(q, year, name, surname).Scan(&category)
	if err != nil {
		log.Printf("Error getting category: %v", (err))
		return category
	}
	return category

}

// GetAvgScoresOfCategories calculates the average scores for
// all the categories in the given year
func (database *DB) GetAvgScoresOfCategories(year int) (avg float64) {
	q := `
SELECT AVG(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
WHERE year = ?
`
	err := database.db.QueryRow(q, year).Scan(&avg)
	if err != nil {
		log.Printf("Error getting average score of categories: %v", err)
		return avg
	}
	return avg
}

// GetAvgScore returns the score's average for the given year and category
func (database *DB) GetAvgScore(year int, category string) (avg float64) {
	q := `
SELECT AVG(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id,
JOIN Giochi G ON G.id = P.giochi
WHERE G.anno = ? AND G.categoria = ?
`
	err := database.db.QueryRow(q, year, category).Scan(&avg)
	if err != nil {
		log.Printf("Error getting average score: %v", err)
		return avg
	}

	return avg
}

// GetMaxScore calculates the maximum score obtained by any player
// in the given year and category
func (database *DB) GetMaxScore(year int, category string) (max float64) {
	q := `
SELECT MAX(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
WHERE G.anno = ? AND G.categoria = ?
`
	err := database.db.QueryRow(q, year, category).Scan(&max)
	if err != nil {
		log.Printf("Error getting max score: %v", err)
		return max
	}
	return max
}

// GetPassword retrieve the hashed password of a user
// password must be saved by going through:
//
// hashPassword := sha256.New()
// hashPassword.Write([]byte("password"))
// return string(hashPassword.Sum(nil))
func (database DB) GetPassword(ctx context.Context, username string) (string, error) {
	var hash string
	q := `SELECT parolachiave FROM Utenti WHERE nomeutente = ?`
	err := database.db.QueryRow(q, username).Scan(&hash)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

// GetReeloCostants retrieve the costants used to calculate the reelo score
func (database *DB) GetReeloCostants() (Costants, error) {
	c := Costants{}
	q := `
SELECT anno_inizio, k_esercizi, finale, fattore_moltiplicativo, exploit, no_partecipazione
FROM Costanti
`

	err := database.db.QueryRow(q).Scan(&c.StartingYear, &c.ExercisesCostant, &c.PFinal, &c.MultiplicativeFactor, &c.AntiExploit, &c.NoPartecipationPenalty)
	if err != nil {
		return c, err
	}
	return c, nil
}

// GetLastKnownYear returns the last year we know anything about
func (database *DB) GetLastKnownYear() (year int) {
	q := `SELECT MAX(anno) FROM Giochi`
	err := database.db.QueryRow(q).Scan(&year)
	if err != nil {
		log.Printf("Error getting last known year: %v", err)
		return year
	}
	return year
}

// GetLastKnownCategoryForPlayer returns the last category into which
// we know the specified player has partecipated
func (database *DB) GetLastKnownCategoryForPlayer(name, surname string) (category string) {
	q := `
SELECT G.categoria FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.cognome = ?
AND G.anno = (
	SELECT MAX(G.anno) FROM Giochi G
	JOIN Partecipazione P ON P.giochi = G.id
	JOIN Giocatore U ON U.id = P.giocatore
	WHERE U.nome = ? AND U.cognome = ?
)
`
	err := database.db.QueryRow(q, name, surname, name, surname).Scan(&category)
	if err != nil {
		log.Printf("Error getting last known category: %v", err)
		return category
	}
	return category
}

// IsResultFromParis checks if the result associated w/ the given data
// comes from an international game
func (database *DB) IsResultFromParis(name, surname string, year int, category string) bool {
	q := `
SELECT P.sede FROM Partecipazione P
JOIN Giocatore U ON U.id = P.giocatore
JOIN Giochi G ON G.id = P.giochi
WHERE U.nome = ? U.cognome = ?
AND G.anno = ? AND G.categoria = ?
`
	var city string
	err := database.db.QueryRow(q, name, surname, year, category).Scan(&city)
	if err != nil {
		log.Printf("Error getting city: %v", err)
		return false
	}
	return city == "paris"
}
