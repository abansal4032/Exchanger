package handlers

import (
	"Exchanger/common"
	"Exchanger/dal"
	"github.com/gorilla/mux"
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