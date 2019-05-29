package db

const (
	findPlayerIDByNameAndSurname = `SELECT id FROM Giocatore WHERE nome = ? AND cognome = ?`
	findPasswordByUsername       = `SELECT parolachiave FROM Utenti WHERE nomeutente = ?`
	findAllPlayers               = `SELECT nome, cognome FROM Giocatore`
	findMaxYear                  = `SELECT MAX(anno) FROM Giochi`

	findAllCostants = `
	SELECT anno_inizio, k_esercizi, finale, fattore_moltiplicativo, exploit, no_partecipazione
	FROM Costanti`

	findResultsByNameAndSurname = `
SELECT R.tempo, R.esercizi, R.punteggio, G.anno, G.categoria
FROM Giocatore U
JOIN Partecipazione P ON P.giocatore = U.id
JOIN Risultato R ON R.id = P.risultato
JOIN Giochi G ON G.id = P.giochi
WHERE U.Nome = ? AND U.Cognome = ?`

	findPartecipationYearsByPlayer = `
SELECT DISTINCT G.anno FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.id = ?`

	findScoresByPlayerAndYear = `
SELECT R.punteggio FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.Cognome = ? AND G.anno = ?`

	findExercisesByPlayerAndYear = `
SELECT R.esercizi FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.Cognome = ? AND G.anno = ?`

	findCategoriesByPlayerAndYear = `
SELECT G.categoria FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.cognome = ? AND G.anno = ? `

	findAvgScoresByYear = `
SELECT AVG(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
WHERE G.anno = ?`

	findAvgScoresByYearAndCategory = `
SELECT AVG(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
WHERE G.anno = ? AND G.categoria = ?`

	findMaxScoreByYearAndCategory = `
SELECT MAX(R.punteggio) FROM Risultato R
JOIN Partecipazione P ON P.risultato = R.id
JOIN Giochi G ON G.id = P.giochi
WHERE G.anno = ? AND G.categoria = ?`

	findCategoryByPlayerAndYear = `
SELECT G.categoria FROM Giochi G
JOIN Partecipazione P ON P.giochi = G.id
JOIN Giocatore U ON U.id = P.giocatore
WHERE U.nome = ? AND U.cognome = ?
AND G.anno = (
	SELECT MAX(G.anno) FROM Giochi G
	JOIN Partecipazione P ON P.giochi = G.id
	JOIN Giocatore U ON U.id = P.giocatore
	WHERE U.nome = ? AND U.cognome = ?
) `

	findCityByPlayerAndYearAndCategory = `
SELECT P.sede FROM Partecipazione P
JOIN Giocatore U ON U.id = P.giocatore
JOIN Giochi G ON G.id = P.giochi
WHERE U.nome = ? AND U.cognome = ?
AND G.anno = ? AND G.categoria = ?`

	findAllPlayersRanks = `
SELECT U.nome, U.cognome, G.categoria, U.reelo FROM Giocatore U
JOIN Partecipazione P ON P.giocatore = U.id
JOIN Giochi G ON G.id = P.giochi
WHERE (G.anno, U.id) IN (
	SELECT MAX(G.anno), U.id FROM Giochi G
	JOIN Partecipazione P ON P.giochi = G.id
	JOIN Giocatore U ON U.id = P.giocatore
	GROUP BY U.id
)`
)

// last played years:
