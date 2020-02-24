package webinterface

import (
	"log"
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Login logs a user in by adding a jwt to the cookies
func (wh *WebserviceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var cred usecases.User

	if err := utils.ReadBody(r.Body, &cred); err != nil {
		wh.Interactor.Log("Cannot read login request body: %v", err)
		http.Error(w, wh.Interactor.Error(err, "E_BAD_BODY", http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	jwt, err := wh.Interactor.Login(cred)
	if err != nil {
		wh.Interactor.Log("Cannot log the user in: %v", err)
		http.Error(w, err.String(), 500)
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
