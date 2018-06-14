package server

import (
	"Exchanger/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io"
	"os"
	"Exchanger/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func accessLogWriteCloser(cfg *Config) io.WriteCloser {
	if cfg.AccessLog == nil {
		return os.Stderr
	}
	if _, err := cfg.AccessLog.Write([]byte("starting server\n")); err != nil {
		log.Printf("unable to log to file %s: %s; defaulting to standard error", cfg.AccessLog.Filename, err.Error())
		return os.Stderr
	}
	return cfg.AccessLog
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")

	router.HandleFunc("/users", handlers.ListUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}/updateToken", handlers.UpdateUserToken).Methods("PATCH")

	router.HandleFunc("/entities", handlers.ListEntities).Methods("GET")
	router.HandleFunc("/entities", handlers.CreateEntity).Methods("POST")
	router.HandleFunc("/entities/{entity_id}", handlers.GetEntity).Methods("GET")
	router.HandleFunc("/entities/search_by_name/{name}", handlers.SearchEntityByName).Methods("GET")
	router.HandleFunc("/entities/search_by_owner/{owner_name}", handlers.GetEntityByOwner).Methods("GET")
	router.HandleFunc("/entities/search_by_requester/{requester_name}", handlers.GetEntityByRequester).Methods("GET")

	router.HandleFunc("/requests", handlers.ListRequests).Methods("GET")
	router.HandleFunc("/requests", handlers.CreateRequest).Methods("POST")
	router.HandleFunc("/requests/{request_id}", handlers.GetRequest).Methods("GET")
	router.HandleFunc("/requests/{request_id}", handlers.UpdateRequest).Methods("PATCH")
	router.HandleFunc("/requests/search_by_requester/{requester_name}", handlers.GetRequestByRequester).Methods("GET")
	router.HandleFunc("/requests/search_by_owner/{owner_name}", handlers.GetRequestByOwner).Methods("GET")
	writeCloser := accessLogWriteCloser(Conf)
	defer writeCloser.Close()
	middleware.AccessLogWriter = writeCloser

	log.Fatal(http.ListenAndServe(":8080", router))
}
