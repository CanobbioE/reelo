package infrastructure

import (
	"encoding/json"
	"fmt"
)

// ErrorHandler implements an error handler utility
type ErrorHandler struct {
}

// NewErrorHandler is the error handler utility constructor
func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

// Error represents an error mesage to be returned to the front end.
// The message should be human readable and the code should quickly
// describe the error.
type Error struct {
	message    string `json:"message"`
	code       string `json:"code"`
	httpStatus int    `json:"status"`
}

// NewError creates a new error
func (eh ErrorHandler) NewError(err error, code string, httpStatus int) fmt.Stringer {
	return Error{
		message:    err.Error(),
		code:       code,
		httpStatus: httpStatus,
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

// Code is the code field getter
func (e Error) Code() string {
	return e.code
}

// Status is the status field getter
func (e Error) Status() int {
	return e.httpStatus
}

// Message is the message field getter
func (e Error) Message() string {
	return e.message
}
