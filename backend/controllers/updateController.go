package controllers

import (
	"context"
	"log"
	"net/http"

	rdb "github.com/CanobbioE/reelo/backend/db"
)

// UpdateDB is to be called from CLI, it is used to automate db updates.
// In production is an empty function
func UpdateDB(w http.ResponseWriter, r *http.Request) {
	log.Println("Called")
	ctx := context.Background()
	db := rdb.Instance()
	ids, err := db.AllPlayersID(ctx)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, id := range ids {
		history, years, err := db.AnalysisHistoryByID(ctx, id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accent := rdb.CreateAccent(history[years[0]][0].Year, 0, history[years[0]][0].City)
		if err = db.UpdateDB(ctx, accent, id); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return
}
