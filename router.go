package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{key}", PutHandler).Methods("PUT")
	router.HandleFunc("/{key}", GetHandler).Methods("GET")
	router.HandleFunc("/{key}", DeleteHandler).Methods("DELETE")

	return router
}
