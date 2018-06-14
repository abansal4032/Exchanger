package handlers

import (
	"net/http"
	"Exchanger/dal"
	"encoding/json"
	"github.com/gorilla/mux"
)

// ListEntities lists the entities
func ListEntities(w http.ResponseWriter, r *http.Request) {
	res, _ := dal.GetAllEntitites("")
	resp, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// GetEntity gets the entities
func GetEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["entity_id"]
	res, _ := dal.GetAllEntitites(id)
	resp, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}