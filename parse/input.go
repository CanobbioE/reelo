package parse

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	rdb "github.com/CanobbioE/reelo/db"
)

const RANK_PATH = "./ranks"

var db *rdb.DB

func init() {
	db = rdb.NewDB()
}

// parseRankingFile reads a ranking from the correct file using the specified
// format. The file's name must be in the format of "year_category.txt"
func parseRankingFile(year int, category, format string) {
	ctx := context.Background()
	filePath := fmt.Sprintf("%s/%d/%d_%s.txt", RANK_PATH, year, year, category)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Couldn't open file.", err)
	}
	defer file.Close()

	// add the current year+category to the db
	gID := db.Add(ctx, "giochi", year, category)

	f := newFormat(strings.Split(strings.ToLower(format), " "))

	// Here we do stuff and things with the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		updateDb(ctx, data, f, gID)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

// updateDb updates the database by adding all the data read from a single row
// of the Ranking's file.
func updateDb(ctx context.Context, data []string, f *Format, gID int) {
	var pID int
	// Add the player information only if he doesn't already exist in the db
	if !db.ContainsPlayer(ctx, data[f.Name], data[f.Surname]) {
		pID = db.Add(ctx, "giocatore", data[f.Name], data[f.Surname])
	} else {
		pID = db.RetrievePlayerID(ctx, data[f.Name], data[f.Surname])
	}
	// Sometime the time is not specified in the ranking file
	time, err := strconv.Atoi(data[f.Time])
	if err != nil || f.Time != -1 {
		time = 0
	}

	rID := db.Add(ctx, "risultato", time, data[f.Exercises], data[f.Score])
	db.Add(ctx, "partecipazione", pID, gID, rID)
}
