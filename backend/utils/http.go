package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
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
