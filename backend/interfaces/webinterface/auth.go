package webinterface

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Login logs a user in by adding a jwt to the cookies
func (wh *WebserviceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var cred usecases.User
	err := utils.ReadBody(r.Body, &cred)
	if err != nil {
		log.Println(err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	status, jwt, err := wh.Interactor.Login(cred)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v", err), status)
		return
	}
	w.Write([]byte(jwt))
	// TODO: this does almost nothing on SPA :(
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: jwt,
	})

	log.Printf("User %s logged in!", cred.Username)
	return
}
