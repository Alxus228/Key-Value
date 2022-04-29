// Package client implements functions that allow us to send http requests to our server.
//
// Available methods: GET all, PUT, DELETE.
package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// String serverAddress contains an actual path to the server.
const serverAddress string = "https://localhost:443"

// newClient returns a link of http.Client, that is able to work with the Https protocol.
func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

// GetAll sends a GET http request without url variables and returns a response from the server.
func GetAll() (*http.Response, error) {
	client := newClient()

	address := serverAddress + "/api/"
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		log.Println("Error during GET all request initialization.")
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}

// Put sends a PUT http request with key and value provided.
func Put(key interface{}, value interface{}) (*http.Response, error) {
	client := newClient()

	json, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	address := serverAddress + "/api/" + fmt.Sprintf("%v", key)
	req, err := http.NewRequest(http.MethodPut, address, bytes.NewBuffer(json))
	if err != nil {
		log.Println("Error during PUT request initialization.")
		log.Println(err)
		return nil, err
	}

	// We set Disable-Logging header as true here, because we generally use
	// client package for benchmark tests or to restore data from "transactions.exe"
	req.Header.Set("Disable-Logging", "true")

	resp, err := client.Do(req)
	return resp, err
}

// Delete sends a DELETE http request with a specific key.
func Delete(key interface{}) (*http.Response, error) {
	client := newClient()

	address := serverAddress + "/api/" + fmt.Sprintf("%v", key)
	req, err := http.NewRequest(http.MethodDelete, address, nil)
	if err != nil {
		log.Println("Error during DELETE request initialization.")
		log.Println(err)
		return nil, err
	}

	// Same here
	req.Header.Set("Disable-Logging", "true")

	resp, err := client.Do(req)
	return resp, err
}
