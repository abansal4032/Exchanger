package handlers

import (
	"Exchanger/common"
	"Exchanger/dal"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
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
	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}