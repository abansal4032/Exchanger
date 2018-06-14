package handlers

import (
	"net/http"
	"Exchanger/models"
	"encoding/json"
	"Exchanger/dal"
	"Exchanger/common"
	"github.com/gorilla/mux"
)

// ListUsers lists the users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dal.GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	common.WriteResponse(w, users)
}

// GetUser gets the user details
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	user, err := dal.GetUsers(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	common.WriteResponse(w, user)
}


// CreateUser creates a new user entity
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	if err := DecodeRequestBody(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while decoding the request body" + err.Error()))
	}
	if err := dal.CreateUser(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// CreateUser creates a new user entity
func UpdateUserToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	req := &models.User{}
	if err := DecodeRequestBody(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while decoding the request body" + err.Error()))
	}
	if err := dal.UpdateUser(req.RegistrationToken, userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func DecodeRequestBody(r *http.Request, model interface{}) error {
	if r.Body == nil {
		return nil
	}
	d := json.NewDecoder(r.Body)
	if err := d.Decode(model); err != nil {
		return err
	}
	return nil
}