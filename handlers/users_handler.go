package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	helpers "github.com/code-sleuth/yummy-recipes-go/helpers"
	"github.com/code-sleuth/yummy-recipes-go/models"
)

// CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "bad request")
	}

	if r.ContentLength == 0 {
		helpers.ErrorResponse(w, http.StatusBadRequest, "empty json body")
	}

	if len(user.Username) < 6 || len(user.Username) > 12 {
		message := "username should be atleast 6 characters and not more than 12 characters"
		helpers.ErrorResponse(w, http.StatusForbidden, message)
	}

	if len(user.Email) == 0 || len(user.Fullname) == 0 || len(user.Password) == 0 {
		helpers.ErrorResponse(w, http.StatusBadRequest, "email, fullname and password should not be empty")
	}

	u, err := models.CreateUser(user.Username, user.Email, user.Fullname, user.Password)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "unable to create user "+err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, u)
}

// GetUsers func
func GetUsers(w http.ResponseWriter, r *http.Request) {

	userList, err := models.GetUsers()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusNoContent, "could not get users from database | "+err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, userList)
}

// GetUser by id
func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	user, err := models.GetUser(params["id"])
	if err != nil {
		helpers.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, user)
}

// UpdateUser func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "bad request")
	}

	u, err := models.UpdateUser(params["id"], user.Email, user.Username, user.Fullname)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, u)
}

// DeleteUser func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	d, err := models.DeleteUser(params["id"])
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, map[string]string{"success": d})
}
