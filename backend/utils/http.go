package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
