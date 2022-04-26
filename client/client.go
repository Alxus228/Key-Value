package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const serverAddress string = "https://localhost:443"

func GetAll() (*http.Response, error) {
	resp, err := http.Get(serverAddress + "/api/")
	return resp, err
}

func Put(key interface{}, value interface{}) (*http.Response, error) {
	client := &http.Client{}

	json, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	address := serverAddress + "/api/" + fmt.Sprintf("%v", key)
	req, err := http.NewRequest(http.MethodPut, address, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error during put request initialization: ", err)
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}

func Delete(key interface{}) (*http.Response, error) {
	client := &http.Client{}

	address := serverAddress + "/api/" + fmt.Sprintf("%v", key)
	req, err := http.NewRequest(http.MethodDelete, address, nil)
	if err != nil {
		log.Fatal("Error during delete request initialization: ", err)
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}
