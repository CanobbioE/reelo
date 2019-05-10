package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/api"
	"github.com/CanobbioE/reelo/backend/services"
)

// Login writes jwt in the HTTP response
// TODO refactoring
func Login(w http.ResponseWriter, r *http.Request) {
	var cred api.Credentials
	err := ReadBody(r.Body, &cred)
	if err != nil {
		log.Println(err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	status, jwt, err := services.Login(cred)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), status)
	}
	w.Write([]byte(jwt))
	log.Printf("User %s logged in!", cred.Username)
	return
}

// ReadBody reads a request body and unmarshall it into a given entity
func ReadBody(r io.Reader, entity interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("Error reading body: %v", err)
	}
	err = json.Unmarshal(body, entity)
	if err != nil {
		return fmt.Errorf("Error unmarshalling body: %v", err)
	}
	return nil
}
