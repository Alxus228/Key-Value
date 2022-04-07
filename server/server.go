package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Run(store *Storage) error {
	router := mux.NewRouter()

	router.HandleFunc("/api/{key}", GetHandler(*store)).Methods("GET")
	router.HandleFunc("/api/", GetAllHandler(*store)).Methods("GET")
	router.HandleFunc("/api/{key}", PutHandler(*store)).Methods("PUT")
	router.HandleFunc("/api/{key}", DeleteHandler(*store)).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	return err
}
