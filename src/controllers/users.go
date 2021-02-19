package controllers

import (
	"api/src/database"
	"api/src/errors"
	"api/src/models"
	"api/src/presentation"
	"api/src/repositories"
	"api/src/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser method to calls user repositories and returns an User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errors.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		errors.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = validations.Prepare(&user); err != nil {
		errors.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	repositorie := repositories.NewUserRepositorie(db)
	user.ID, err = repositorie.Create(user)
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	presentation.JSON(w, http.StatusCreated, user)
}

// ListUsers List all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	users, err := repositorie.List(nameOrNick)
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	presentation.JSON(w, http.StatusOK, users)
}

// GetUser Find one user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		errors.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	user, err := repositorie.GetUserByID(userID)
	if err != nil {
		errors.Err(w, http.StatusInternalServerError, err)
		return
	}

	presentation.JSON(w, http.StatusOK, user)
}

// UpdateUser update one user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeleteUser delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
