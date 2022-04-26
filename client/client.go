package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const serverAddress string = "https://localhost:443"

func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

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

	resp, err := client.Do(req)
	return resp, err
}

func Delete(key interface{}) (*http.Response, error) {
	client := newClient()

	address := serverAddress + "/api/" + fmt.Sprintf("%v", key)
	req, err := http.NewRequest(http.MethodDelete, address, nil)
	if err != nil {
		log.Println("Error during DELETE request initialization.")
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}
