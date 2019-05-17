package services

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/CanobbioE/reelo/backend/api"
	rdb "github.com/CanobbioE/reelo/backend/db"
	"github.com/CanobbioE/reelo/backend/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// Login implements the login logic.
// It returns an http status and an eventual error.
func Login(c api.Credentials) (int, string, error) {
	var jwt string

	db := rdb.NewDB()
	expPassword, err := db.Password(context.Background(), c.Username)
	defer db.Close()

	if err != nil {
		// TODO: can't compare this error
		if err == fmt.Errorf("user not found") {
			return http.StatusUnauthorized, jwt, fmt.Errorf("user not found: %v", err)
		}
		return http.StatusUnauthorized, jwt, fmt.Errorf("error while reading from db: %v", err)
	}

	if toHexHash(c.Password) != expPassword {
		return http.StatusUnauthorized, jwt, fmt.Errorf("passwords don't match")
	}

	jwt, err = generateJWT(c.Username)
	if err != nil {
		return http.StatusInternalServerError, jwt, fmt.Errorf("error while generating the jwt %v", err)
	}

	return http.StatusOK, jwt, nil
}

func toHexHash(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", string(hash.Sum(nil)))
}

func generateJWT(username string) (string, error) {
	jwtKey := utils.JWTKey()

	// Declare the expiration time of the token
	// here, we have kept it as 60 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	c := &utils.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
