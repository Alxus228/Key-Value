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
			log.Println("GET request with no key.")
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		value, err := s.Get(key)
		if err != nil {
			log.Println("Item", key, "hasn't been found during GET request.")
			log.Println(err)
			http.Error(writer, "such key doesn't exist", http.StatusNotFound)
			return
		}

		response, serialErr := json.Marshal(value)
		if serialErr != nil {
			log.Println("Serialization failed:", serialErr)
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
			log.Println("GET all request onto an empty storage.")
			http.Error(writer, "storage is empty", http.StatusNotFound)
			return
		}

		t := template.Must(template.ParseFiles("server/get_all.tmpl"))

		err := t.Execute(writer, data)
		if err != nil {
			log.Println("Internal error during get all request.")
			log.Println(err)
			http.Error(writer, "something went wrong", http.StatusInternalServerError)
		}

		log.Println("Successfully returned storage to user.")
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
			log.Println("PUT request with no key.")
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		var value interface{}
		deserialErr := json.NewDecoder(request.Body).Decode(&value)
		if deserialErr != nil {
			log.Println("Deserialization failed:", deserialErr)
			http.Error(writer, "data is unserializable", http.StatusBadRequest)
			return
		}

		err := s.Put(key, value)
		if err != nil {
			log.Println("Internal error during PUT request.")
			log.Println("Key:", key, "value:", value, "error:", err)
			http.Error(writer, "hasn't succeeded to save the value", http.StatusInternalServerError)
			return
		}

		loggingHeader := request.Header.Get("Disable-Logging")
		if loggingHeader != "true" {
			logTransaction("PUT", key, value)
		}

		log.Println("item", key, "successfully created.")
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
			log.Println("DELETE request with no key.")
			http.Error(writer, "key is empty in the URL", http.StatusBadRequest)
			return
		}

		_, getErr := s.Get(key)
		if getErr != nil {
			log.Println("Item", key, "hasn't been found during DELETE request.")
			log.Println(getErr)
			http.Error(writer, "such key doesn't exist", http.StatusNotFound)
			return
		}

		err := s.Delete(key)
		if err != nil {
			log.Println("Internal error during DELETE request.")
			log.Println("Key:", key, "error:", err)
			http.Error(writer, "hasn't succeeded to delete the key", http.StatusInternalServerError)
			return
		}

		loggingHeader := request.Header.Get("Disable-Logging")
		if loggingHeader != "true" {
			logTransaction("DELETE", key, nil)
		}

		log.Println("item ", key, " successfully deleted.")
		writer.WriteHeader(http.StatusNoContent)
	}
}
