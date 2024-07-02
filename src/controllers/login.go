package controllers

import (
	"api-rede-social/src/authentication"
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"api-rede-social/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login make users login
func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(bodyRequest, &user); err != nil {
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
	DbUser, err := repositorie.FindByEmail(user.Email)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(user.Password, DbUser.Password); err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(DbUser.ID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}
