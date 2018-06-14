package server

import (
	"Exchanger/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/users", handlers.ListUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}/updateToken", handlers.UpdateUserToken).Methods("PATCH")
	router.HandleFunc("/entities", handlers.ListEntities).Methods("GET")
	router.HandleFunc("/entities/{entity_id}", handlers.GetEntity).Methods("GET")
	router.HandleFunc("/requests", handlers.ListRequests).Methods("GET")
	router.HandleFunc("/requests/{request_id}", handlers.GetRequest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
