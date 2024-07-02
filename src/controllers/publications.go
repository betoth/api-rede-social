package controllers

import (
	"api-rede-social/src/authentication"
	"api-rede-social/src/database"
	"api-rede-social/src/models"
	"api-rede-social/src/repositories"
	"api-rede-social/src/response"
	"encoding/json"
	"io"
	"net/http"
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
	w.Write([]byte("Find"))

}

// FindPublicationByID find a publication by ID
func FindPublicationByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindByID"))

}

// UpdatePublication by id
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))

}

// DeletePublication delete a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))

}
