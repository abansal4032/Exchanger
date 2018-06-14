package handlers

import (
	"net/http"
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
