package dto

import "github.com/CanobbioE/reelo/backend/domain"

// FileUpload represents the information associated w/
// an uploaded ranking file.
type FileUpload struct {
	Game   domain.Game `json:"game"`
	Format string      `json:"format"`
}

// Error represents an error mesage to be returned to the front end.
// The message should be human readable and the code should quickly
// describe the error.
type Error struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	HTTPStatus int    `json:"status"`
}
