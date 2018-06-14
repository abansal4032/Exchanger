package server

import (
	"fmt"
	"net/http"
	"log"
	"Exchanger/handlers"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/users", handlers.ListUsers).Methods("GET")
	router.HandleFunc("/users/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/entities", handlers.ListEntities).Methods("GET")
	router.HandleFunc("/entities/{entity_id}", handlers.GetEntity).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
