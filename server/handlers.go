package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// Storage interface describes 4 methods: Get, GetAll, Put and Delete
//
// We use this interface, instead of storage type from storage package,
// because our application use principalities of cloud-native development,
// and we might use other storage, not the local one.
type Storage interface {
	// Get receives a key, and returns a value.
	Get(interface{}) (interface{}, error)
	// GetAll returns key-value map.
	GetAll() (map[interface{}]interface{}, error)
	// Put receives a key and a value, and creates, or modifies a key-value pair in the storage.
	Put(interface{}, interface{}) error
	// Delete receives a key, and dispose of the corresponding key-value pair.
	Delete(interface{}) error
}

// getHandler responses with:
//   * 200 code and body with an item from key-value storage, when the key-value pair is presented in the storage.
//   * 400 code, when the key isn't presented.
//   * 404 code, when there is no element with such key in the storage.
//   * 500 code, when something when wrong, and it hasn't managed to serialize the item from the storage.
func getHandler(s Storage) http.HandlerFunc {
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

// getAllHandler responses with:
//   * 200 code and a html representation of key-value map, in case storage has something in it.
//   * 400 code, when the storage is empty.
func getAllHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, getAllErr := s.GetAll()
		if getAllErr != nil {
			http.Error(writer, "storage is empty", http.StatusNotFound)
			return
		}

		t := template.Must(template.ParseFiles("server/get_all.tmpl"))

		err := t.Execute(writer, data)
		if err != nil {
			log.Println("Internal error during get all request: ", err)
			http.Error(writer, "something went wrong", http.StatusInternalServerError)
		}
	}
}

// putHandler responses with:
//   * 201 code, after a successful item creation.
//   * 400 code, when the key isn't presented or data cannot be deserialized.
//   * 500 code, when something when wrong, and it hasn't managed to create the item.
func putHandler(s Storage) http.HandlerFunc {
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
			log.Println("Internal error during put request: ", err)
			http.Error(writer, "hasn't succeeded to save the value", http.StatusInternalServerError)
			return
		}

		logTransaction("PUT", key, value)
		writer.WriteHeader(http.StatusCreated)
	}
}

// deleteHandler responses with:
//   * 204 code, after successful deletion.
//   * 400 code, when the key isn't presented.
//   * 404 code, when the item already doesn't exist.
//   * 500 code, when something when wrong, and it hasn't managed to delete the item.
func deleteHandler(s Storage) http.HandlerFunc {
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
		if err != nil {
			log.Println("Internal error during delete request: ", err)
			http.Error(writer, "hasn't succeeded to delete the key", http.StatusInternalServerError)
			return
		}

		logTransaction("DELETE", key, nil)
		writer.WriteHeader(http.StatusNoContent)
	}
}
