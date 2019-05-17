package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// ReadBody reads a request body and unmarshall it into a given entity
func ReadBody(r io.Reader, entity interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("Error reading body: %v", err)
	}
	err = json.Unmarshal(body, entity)
	if err != nil {
		return fmt.Errorf("Error unmarshalling body: %v", err)
	}
	return nil
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
