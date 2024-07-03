package controllers

import (
	"api-rede-social/src/authentication"
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication create a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var publication models.Publication

	if err := json.Unmarshal(bodyRequest, &publication); err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	publication.AuthorID, err = authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusForbidden, err)
		return
	}

	if err := publication.Prepare(); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	repositorie := repositories.NewPublicationRepositorie(db)
	publication.ID, err = repositorie.Create(publication, publication.AuthorID)

	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

// FindPublications list all relacioned publications
func FindPublications(w http.ResponseWriter, r *http.Request) {

	userIDToken, err := authentication.ExtractUserID(r)
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

	PubRepositorie := repositories.NewPublicationRepositorie(db)
	publications, err := PubRepositorie.FindPublications(userIDToken)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

// FindPublicationByID find a publication by ID
func FindPublicationByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repositorie := repositories.NewPublicationRepositorie(db)
	publication, err := repositorie.FindPublicationByID(publicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

// UpdatePublication by id
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	PublicationID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDTOken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusForbidden, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	repositorie := repositories.NewPublicationRepositorie(db)
	publicationDB, err := repositorie.FindPublicationByID(PublicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if userIDTOken != publicationDB.AuthorID {
		response.ErrorJSON(w, http.StatusForbidden, errors.New("Forbidden"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	err = json.Unmarshal(requestBody, &publication)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = repositorie.UpdatePublication(PublicationID, publication)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePublication delete a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	PublicationID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDTOken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.ErrorJSON(w, http.StatusForbidden, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	repositorie := repositories.NewPublicationRepositorie(db)
	publicationDB, err := repositorie.FindPublicationByID(PublicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if userIDTOken != publicationDB.AuthorID {
		response.ErrorJSON(w, http.StatusForbidden, errors.New("Forbidden"))
		return
	}

	err = repositorie.DeletePublication(PublicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, err)

}

// FindPublicationByUser list all publications of an user
func FindPublicationByUser(w http.ResponseWriter, r *http.Request) {

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

	repositorie := repositories.NewPublicationRepositorie(db)
	publications, err := repositorie.FindPublicationByUser(userID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

// LikePublication like a publication
func LikePublication(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameter["id"], 10, 64)
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

	repositorie := repositories.NewPublicationRepositorie(db)
	err = repositorie.LikePublication(publicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnlikePublication unlike a publication
func UnlikePublication(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameter["id"], 10, 64)
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

	repositorie := repositories.NewPublicationRepositorie(db)
	err = repositorie.UnlikePublication(publicationID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusNoContent, nil)

}
