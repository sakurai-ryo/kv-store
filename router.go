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
			"Length": r.ContentLength,
			"Ua":     r.UserAgent,
		}).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

func maxReqBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, MAX_BODY_SIZE)
		next.ServeHTTP(w, r)
	})
}

func NewRouter(handler StoreHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.Use(maxReqBodyMiddleware)

	router.HandleFunc("/{key}", handler.PutHandler).Methods("PUT")
	router.HandleFunc("/{key}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/{key}", handler.DeleteHandler).Methods("DELETE")

	return router
}
