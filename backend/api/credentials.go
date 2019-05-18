package api

// Credentials represent a user login credentials
type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}
