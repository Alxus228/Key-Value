package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Run(store *Storage) error {
	router := mux.NewRouter()

	router.HandleFunc("/api/", GetAllHandler(*store)).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	return err
}
