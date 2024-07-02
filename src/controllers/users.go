package controllers

import (
	"api-rede-social/src/authentication"
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"api-rede-social/src/security"
	"encoding/json"
	"errors"
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
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	userIDTOken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return

	}

	if userID != userIDTOken {

		response.ErrorJSON(w, http.StatusForbidden, errors.New("Forbidden"))
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

	if err = repositori.UpdateUser(userID, user); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delete a user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDTOken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return

	}

	if userID != userIDTOken {

		response.ErrorJSON(w, http.StatusForbidden, errors.New("Forbidden"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	if err = repositorie.Delete(userID); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// FollowUser follow a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	UserID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if followerID == UserID {
		response.ErrorJSON(w, http.StatusForbidden, errors.New("It is not possible to follow the same user"))
		return
	}

	repositorie := repositories.NewUserRepositorie(db)
	if err := repositorie.Follow(UserID, followerID); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollow a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	followID, err := authentication.ExtractUserID(r)

	if userID == followID {

		response.ErrorJSON(w, http.StatusForbidden, errors.New("It is not possible to unfollow the same user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)

	if err := repositorie.Unfollow(userID, followID); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// SearchFollowers list all followers
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["id"], 10, 64)
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
	users, err := repositorie.SearchFollowers(userID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)

}

// SearchFollowing list all users if follows a user
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)

	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	users, err := repositorie.SearchFollowing(userID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)

}

// UpdatePassword update user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDTOken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDTOken {
		response.ErrorJSON(w, http.StatusUnauthorized, errors.New("Forbidden"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var password models.Password

	if err := json.Unmarshal(bodyRequest, &password); err != nil {
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
	passwordFromDb, err := repositorie.SearchPasswordByID(userID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(password.CurrentPassword, passwordFromDb); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	passwordHash, err := security.Hash(password.NewPassword)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	if err := repositorie.UpdatePassword(userID, string(passwordHash)); err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
