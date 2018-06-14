package handlers

import (
	"Exchanger/common"
	"Exchanger/dal"
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
