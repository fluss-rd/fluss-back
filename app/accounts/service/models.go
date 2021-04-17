package service

// LoginResponse login response
type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"userID"`
}
