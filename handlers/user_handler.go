package handlers

import "net/http"

// ListUsers lists the users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Users"))
}