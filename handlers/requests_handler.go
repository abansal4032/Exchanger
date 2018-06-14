package handlers

import (
	"Exchanger/common"
	"Exchanger/dal"
	"github.com/gorilla/mux"
	"Exchanger/models"
	"net/http"
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

func GetRequestByRequester(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requesterName := vars["requester_name"]
	requests, err := dal.GetRequestsByRequester(requesterName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, requests)
}

func GetRequestByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ownerName := vars["owner_name"]
	requests, err := dal.GetRequestsByOwner(ownerName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, requests)
}

// CreateUser creates a new user entity
func CreateRequest(w http.ResponseWriter, r *http.Request) {
	req := &models.Requests{}
	if err := DecodeRequestBody(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while decoding the request body" + err.Error()))
	}
	if err := dal.CreateRequest(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// CreateUser creates a new user entity
func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["request_id"]
	req := &models.Requests{}
	if err := DecodeRequestBody(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while decoding the request body" + err.Error()))
	}
	if err := dal.UpdateRequest(req, requestId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
