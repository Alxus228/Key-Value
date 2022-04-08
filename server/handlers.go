package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Storage interface {
	Get(interface{}) (interface{}, error)
	GetAll() (map[interface{}]interface{}, error)
	Put(interface{}, interface{}) error
	Delete(interface{}) error
}

func GetHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		urlVariables := mux.Vars(request)

		key, ok := urlVariables["key"]
		if !ok {
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		value, err := s.Get(key)
		if err != nil {
			http.Error(writer, "such key doesn't exist", http.StatusNotFound)
			return
		}

		response, serialErr := json.Marshal(value)
		if serialErr != nil {
			http.Error(writer, "serialization error", http.StatusInternalServerError)
			return
		}

		writer.Write(response)
	}
}

func GetAllHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, getAllErr := s.GetAll()
		if getAllErr != nil {
			http.Error(writer, "storage is empty", http.StatusNotFound)
			return
		}

		t := template.Must(template.ParseFiles("server/get_all.tmpl"))

		err := t.Execute(writer, data)
		if err != nil {
			http.Error(writer, "something went wrong", http.StatusInternalServerError)
		}
	}
}

func PutHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		urlVariables := mux.Vars(request)

		key, ok := urlVariables["key"]
		if !ok {
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		var value interface{}
		deserialErr := json.NewDecoder(request.Body).Decode(&value)
		if deserialErr != nil {
			http.Error(writer, "data is unserializable", http.StatusBadRequest)
			return
		}

		err := s.Put(key, value)
		if err != nil {
			http.Error(writer, "hasn't succeeded to save the value", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	}
}

func DeleteHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		urlVariables := mux.Vars(request)

		key, ok := urlVariables["key"]
		if !ok {
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		_, getErr := s.Get(key)
		if getErr != nil {
			http.Error(writer, "such key doesn't exist", http.StatusNotFound)
			return
		}

		err := s.Delete(key)
		if err == nil {
			http.Error(writer, "hasn't succeeded to delete the key", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}
