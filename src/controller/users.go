package controller

import (
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("create"); err != nil {
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
	user, err := repositorie.Find(searchCondition)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// SearchUser searches for a specific user in database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
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

	user, err := repositorie.FindByID(ID)
	if err != nil {

		if err.Error() == "User not found" {
			response.ErrorJSON(w, http.StatusNotFound, err)
		} else {
			response.ErrorJSON(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// UpdateUser update a user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositori := repositories.NewUserRepositorie(db)

	if err = repositori.UpdateUser(ID, user); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delete a user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted"))
}
