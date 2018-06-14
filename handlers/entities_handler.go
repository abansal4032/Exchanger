package handlers

import (
	"Exchanger/common"
	"Exchanger/dal"
	"Exchanger/models"
	"github.com/gorilla/mux"
	"net/http"
)

// ListEntities lists the entities
func ListEntities(w http.ResponseWriter, r *http.Request) {
	res, err := dal.GetAllEntitites("")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, res)
}

// GetEntity gets the entity by id
func GetEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["entity_id"]
	res, err := dal.GetAllEntitites(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, res[0])
}

// CreateUser creates a new user entity
func CreateEntity(w http.ResponseWriter, r *http.Request) {
	req := &models.Entity{}
	if err := DecodeRequestBody(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while decoding the request body" + err.Error()))
	}
	if err := dal.CreateEntity(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// SearchEntityByName searches the entities for the given string name
func SearchEntityByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchString := vars["name"]
	res, err := dal.SearchEntititesByName(searchString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, res)
}

// GetEntityByOwner searches the entities for the given owner name
func GetEntityByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner_name"]
	res, err := dal.GetEntityByOwner(owner)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, res)
}

// GetEntityByRequester searches the entities for the given requester name
func GetEntityByRequester(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requester := vars["requester_name"]
	res, err := dal.GetEntityByRequester(requester)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	common.WriteResponse(w, res)
}