package dto

import (
	"github.com/CanobbioE/reelo/backend/domain"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/dgrijalva/jwt-go"
)

// FileUpload represents the information associated w/
// an uploaded ranking file.
type FileUpload struct {
	Game   domain.Game `json:"game"`
	Format string      `json:"format"`
	City   string      `json:"city"`
}

// Error represents an error mesage to be returned to the front end.
// The message should be human readable and the code should quickly
// describe the error.
type Error struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	HTTPStatus int    `json:"status"`
}

// SlimPartecipation represents a simplified partecipation relationship
type SlimPartecipation struct {
	City         string  `json:"city"`
	Category     string  `json:"category"`
	IsParis      bool    `json:"isParis"`
	Year         int     `json:"year"`
	MaxExercises int     `json:"eMax"`
	MaxScore     int     `json:"dMax"`
	Score        int     `json:"d"`
	Exercises    int     `json:"e"`
	Time         int     `json:"time"`
	Position     int     `json:"position"`
	PseudoReelo  float64 `json:"pseudoReelo"`
}

// History is a collection of simplified partecipations
type History []SlimPartecipation

// HistoryByYear is an history indexed by parteciaption year to simplify resarch
type HistoryByYear map[int]History

// Claims is a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Rank struct {
	LastCategory string                           `json:"lastCategory"`
	Player       domain.Player                    `json:"player"`
	History      usecases.SlimPartecipationByYear `json:"history"`
}
