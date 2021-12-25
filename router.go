package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"host":   r.Host,
			"path":   r.URL.Path,
			"Ua":     r.UserAgent,
		}).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	router.HandleFunc("/{key}", PutHandler).Methods("PUT")
	router.HandleFunc("/{key}", GetHandler).Methods("GET")
	router.HandleFunc("/{key}", DeleteHandler).Methods("DELETE")

	return router
}
