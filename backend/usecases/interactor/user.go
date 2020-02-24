package interactor

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CanobbioE/reelo/backend/interfaces/webinterface/dto"
	"github.com/CanobbioE/reelo/backend/usecases"
	jwt "github.com/dgrijalva/jwt-go"
)

// Login implements the login logic.
// It returns an http status, the jwt and an eventual error.
func (i *Interactor) Login(user usecases.User) (string, fmt.Stringer) {
	var jwt string

	expPassword, err := i.UserRepository.FindPasswordByUsername(context.Background(), user.Username)
	if err != nil {
		// TODO: can't compare this error
		if err.Error() == "user not found" {
			return jwt, i.ErrorHandler.NewError(err, "E_NO_AUTH", http.StatusUnauthorized)
		}
		return jwt, i.ErrorHandler.NewError(err, "E_UNEXPECTED", http.StatusInternalServerError)
	}

	if toHexHash(user.Password) != expPassword {
		return jwt, i.ErrorHandler.NewError(fmt.Errorf("passwords don't match"), "E_NO_AUTH", http.StatusUnauthorized)
	}

	jwt, err = generateJWT(user.Username)
	if err != nil {
		return jwt, i.ErrorHandler.NewError(err, "E_UNEXPECTED", http.StatusInternalServerError)
	}

	return jwt, nil
}

func toHexHash(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", string(hash.Sum(nil)))
}

func generateJWT(username string) (string, error) {
	jwtKey := JWTKey()

	// Declare the expiration time of the token
	// here, we have kept it as 3h
	expirationTime := time.Now().Add(180 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	c := &dto.Claims{
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

// JWTKey creates the JWT key used to create the signature
// using either the hardcoded dev string or the PROD environment variable
func JWTKey() []byte {
	k := os.Getenv("JWT_KEY")
	if k == "" {
		k = "my_secret_key"
	}

	return []byte("my_secret_key")
}
