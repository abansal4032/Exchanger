package server

import (
	"fmt"
	"net/http"
	"log"
	"Exchanger/server/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Start() {
	registerRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerRoutes() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/users", handlers.ListUsers)
}