package httphandler

// LoginRequestBody what the login request body should look like
type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
