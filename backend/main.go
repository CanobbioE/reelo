package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/upload", CreateRankingFile).Methods("POST")
	router.HandleFunc("/ranks", GetRanks).Methods("GET")
	router.HandleFunc("/admin", Login).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}

// Login writes jwt in the HTTP response
// TODO refactoring
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	type cred struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}
	var c cred

	// Create the JWT key used to create the signature
	var jwtKey = []byte("TODO")

	// Create a struct that will be encoded to a JWT.
	// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
	type claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		http.Error(w, "can't unmarshall body", http.StatusBadRequest)
		return
	}

	db := rdb.NewDB()

	expPassword, err := db.GetPassword(context.Background(), c.Username)
	if err != nil {
		// TODO: can't compare this error
		if err == fmt.Errorf("user not found") {
			log.Printf("Error user not found: %v", err)
			http.Error(w, "can't login", http.StatusUnauthorized)
			return
		}
		log.Printf("Error reading from db: %v", err)
		http.Error(w, "can't login", http.StatusUnauthorized)
		return
	}
	hashPassword := sha256.New()
	hashPassword.Write([]byte(c.Password))

	if string(hashPassword.Sum(nil)) == expPassword {
		// Declare the expiration time of the token
		// here, we have kept it as 60 minutes
		expirationTime := time.Now().Add(60 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &claims{
			Username: c.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			http.Error(w, "can't login", http.StatusUnauthorized)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		log.Println("Logged in!")
		return
	}
	http.Error(w, "can't login", http.StatusUnauthorized)
	return
}

// Stuff does things
func Stuff(w http.ResponseWriter, r *http.Request) {
	// db := rdb.NewDB()
	// dataAll := parse.All()
	// for year, lines := range dataAll {
	//    for _, line := range lines {
	//       playerID := db.Add(context.Background(), "giocatore", line.name, line.surname, 0)
	//       resultID := db.Add(context.Background(), "risultato", line.tempo, line.esercizi, line.punteggio)
	//       gameID := db.Add(context.Background(), "giochi", year, line.categoria)
	//       db.Add(context.Background(), "partecipazione", playerID, gameID, resultID, line.sede)

	//    }
	// }

	// TODO: import to db
}

// GetRanks returns a list of all the ranks in the database
func GetRanks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}

// CreateRankingFile creates a new ranking file
// TODO: authentication
func CreateRankingFile(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// var person Person
	// _ = json.NewDecoder(r.Body).Decode(&person)
	// person.ID = params["id"]
	// people = append(people, person)
	// json.NewEncoder(w).Encode(people)
}
