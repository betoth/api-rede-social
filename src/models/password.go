package models

// Password struct represents format of requitition for password alteration
type Password struct {
	NewPassword     string `json:"new-password"`
	CurrentPassword string `json:"current-password"`
}
