package webinterface

import (
	"net/http"

	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
)

// Login logs a user in by adding a jwt to the cookies
func (wh *WebserviceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var cred usecases.User
	if err := utils.ReadBody(r.Body, &cred); err != nil {
		wh.Interactor.Log("Cannot read login request body: %v", err)
		http.Error(w, utils.NewError(err, "E_BAD_REQ", http.StatusBadRequest).String(), http.StatusBadRequest)
		return
	}
	jwt, err := wh.Interactor.Login(cred)
	if !err.IsNil {
		http.Error(w, err.String(), err.HTTPStatus)
		return
	}
	w.Write([]byte(jwt))
	// TODO: this does almost nothing on SPA :(
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: jwt,
	})

	wh.Interactor.Log("User %s logged in!", cred.Username)
	return
}
