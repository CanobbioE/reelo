package utils

import (
	"encoding/json"
)

// Error represents an error mesage to be returned to the front end.
// The message should be human readable and the code should quickly
// describe the error.
type Error struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	HTTPStatus int    `json:"status"`
	IsNil      bool   `json:"-"`
}

// NewError creates a new error
func NewError(err error, code string, httpStatus int) Error {
	if err == nil {
		return NewNilError()
	}
	return Error{
		Message:    err.Error(),
		Code:       code,
		HTTPStatus: httpStatus,
		IsNil:      false,
	}
}

// String marshals the given dto Error into an http retunable string
// implementing the Stringer interface
func (e Error) String() string {
	ret, err := json.Marshal(e)
	if err != nil {
		// I believe we are allowed to panic if we fail to marshall an error
		// to send to the front end. This error will only be raised
		// by programming mistakes
		panic(err)
	}
	return string(ret)
}

func NewNilError() Error {
	return Error{
		IsNil: true,
	}
}
