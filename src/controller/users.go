package controller

import (
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// CreateUser create a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	user.ID, err = repositorie.Create(user)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
	return

}

// SearchUsers searches all users in database
func SearchUsers(w http.ResponseWriter, r *http.Request) {

	searchCondition := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	user, err := repositorie.Search(searchCondition)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if len(user) == 0 {
		err = errors.New("No users found")
		response.ErrorJSON(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, user)
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
