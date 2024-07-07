package models

// AuthenticationData represents response from authentication
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
