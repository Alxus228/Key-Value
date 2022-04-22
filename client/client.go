package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const serverAddress string = "http://localhost:8080"

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
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}
