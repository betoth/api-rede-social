package controller

import "net/http"

// CreateUser create a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}

// SearchUsers searches all users in database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User locate"))
}

// SearchUser searches for a specific user in database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("localized users"))
}

// UpdateUser update a user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User updated"))
}

// DeleteUser delete a user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted"))
}
