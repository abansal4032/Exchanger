package handlers

import (
	"net/http"
	"Exchanger/dal"
	"Exchanger/common"
	"github.com/gorilla/mux"
)

// ListUsers lists the users
func ListRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := dal.GetRequests("")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, requests)
}

// GetUser gets the user details
func GetRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["request_id"]
	request, err := dal.GetRequests(requestId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, request[0])
}

