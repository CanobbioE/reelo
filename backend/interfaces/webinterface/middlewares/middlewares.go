package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/CanobbioE/reelo/backend/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// RequireAuth is a middleware that checks the user authorization to make a request
// TODO: we are passing jtw in headers and bodies, we should just use cookies
func RequireAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")
		if token == "" || token == "null" {
			log.Println("Missing token")
			http.Error(w, utils.NewError(fmt.Errorf("missing token"), "E_NO_AUTH", http.StatusUnauthorized).String(), http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims,
			func(token *jwt.Token) (interface{}, error) {
				return JWTKey(), nil
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

// Claims is a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTKey creates the JWT key used to create the signature
// using either the hardcoded dev string or the PROD environment variable
func JWTKey() []byte {
	k := os.Getenv("JWT_KEY")
	if k == "" {
		k = "my_secret_key"
	}

	return []byte("my_secret_key")
}
