package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadGateway, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all Users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
