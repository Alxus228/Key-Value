// Package server implements a function, that runs a http server
// and handles queries, that modify the local storage.
//
// Pass a link of the variable, which type describes Storage
// interface, to Run function in order to set up the server.
package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Run creates a new router that listens to the port 8080, and handles requests from the client.
func Run(store *Storage) error {
	router := mux.NewRouter()

	router.HandleFunc("/api/{key}", getHandler(*store)).Methods("GET")
	router.HandleFunc("/api/", getAllHandler(*store)).Methods("GET")
	router.HandleFunc("/api/{key}", putHandler(*store)).Methods("PUT")
	router.HandleFunc("/api/{key}", deleteHandler(*store)).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	return err
}
