package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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

// Paginate extrapolate the page's number and size from the given http request
func Paginate(r *http.Request) (page, size int, err error) {
	pageString := r.URL.Query().Get("page")
	sizeString := r.URL.Query().Get("size")

	page, err = strconv.Atoi(string(pageString))
	if err != nil {
		return page, size, fmt.Errorf("error converting page: %v", err)
	}

	size, err = strconv.Atoi(string(sizeString))
	if err != nil {
		return page, size, fmt.Errorf("error converting size: %v", err)
	}

	return page, size, nil
}
