package main

import "net/http"

type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
	//Err() <-chan error
	//ReadEvents() (<-chan Event, <-chan error)
	//Run()
}

type StoreHandler interface {
	PutHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
}
