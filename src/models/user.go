package models

import (
	"errors"
	"strings"
	"time"
)

// User struct represent users
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nick_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare apply validations and format user input
func (user *User) Prepare() error {

	if err := user.Validate(); err != nil {
		return err
	}

	user.Format()
	return nil
}

// Validate apply validations in user input
func (user *User) Validate() error {

	if user.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if user.NickName == "" {
		return errors.New("Nick Name cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("Password cannot be empty")
	}

	return nil
}

// Format remove spaces from user input
func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
}
