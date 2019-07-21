package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/CanobbioE/reelo/backend/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// Auth is a middleware that checks the user authorization to make a request
// TODO: we are passing jtw in headers and bodies, we should just use cookies
func Auth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")
		if token == "" || token == "null" {
			log.Println("Missing token")
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		claims := &utils.Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims,
			func(token *jwt.Token) (interface{}, error) {
				return utils.JWTKey(), nil
			})
		if !tkn.Valid {
			log.Println("Invalid token")
			http.Error(w, "Autenticazione non valida - esci dall'applicazione e autenticati nuovamente", http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {

				log.Println("Invalid signature")
				http.Error(w, "invalid signature", http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
