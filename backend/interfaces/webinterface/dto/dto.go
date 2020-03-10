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

// SlimParticipation represents a simplified participation relationship
type SlimParticipation struct {
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

// History is a collection of simplified participations
type History []SlimParticipation

// HistoryByYear is an history indexed by participation year to simplify resarch
type HistoryByYear map[int]History

// Claims is a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Rank represents a row on the ranking list
type Rank struct {
	LastCategory string                           `json:"lastCategory"`
	Player       domain.Player                    `json:"player"`
	History      usecases.SlimParticipationByYear `json:"history"`
}
