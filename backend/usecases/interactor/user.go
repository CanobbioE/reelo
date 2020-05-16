package interactor

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/CanobbioE/reelo/backend/interfaces/webinterface/dto"
	"github.com/CanobbioE/reelo/backend/usecases"
	"github.com/CanobbioE/reelo/backend/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// Login implements the login logic.
// It returns an http status, the jwt and an eventual error.
// The function is called when a User tries to access the application's
// administration panel.
func (i *Interactor) Login(user usecases.User) (string, utils.Error) {
	var jwt string

	expPassword, err := i.UserRepository.FindPasswordByUsername(context.Background(), user.Username)
	if err != nil {
		// TODO: can't compare this error
		if err.Error() == "no values in result set" {
			i.Logger.Log("Login: cannot find user: %v", err)
			return jwt, utils.NewError(fmt.Errorf("user does not exist"), "E_NO_AUTH", 401)
		}
		i.Logger.Log("Login: cannot find password: %v", err)
		return jwt, utils.NewError(err, "E_GENERIC", 500)
	}

	if toHexHash(user.Password) != expPassword {
		i.Logger.Log("Login: wrong password")
		return jwt, utils.NewError(fmt.Errorf("wrong password"), "E_NO_AUTH", 401)
	}

	jwt, err = generateJWT(user.Username)
	if err != nil {
		i.Logger.Log("Login: cannot generate JWT: %v", err)
		return jwt, utils.NewError(err, "E_GENERIC", 500)
	}

	return jwt, utils.NewNilError()
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
