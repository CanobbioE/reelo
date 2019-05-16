package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/api"
	"github.com/CanobbioE/reelo/backend/services"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Login writes jwt in the HTTP response
func Login(w http.ResponseWriter, r *http.Request) {
	var cred api.Credentials
	err := utils.ReadBody(r.Body, &cred)
	if err != nil {
		log.Println(err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	status, jwt, err := services.Login(cred)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v", err), status)
		return
	}
	w.Write([]byte(jwt))
	log.Printf("User %s logged in!", cred.Username)
	return
}
