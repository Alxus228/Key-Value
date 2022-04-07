package server

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Storage interface {
	Get(interface{}) (interface{}, error)
	GetAll() map[interface{}]interface{}
	Put(interface{}, interface{})
	Delete(interface{})
}

func GetAllHandler(s Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := s.GetAll()

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
		}

		s.Put(key, request.Body)
		_, err := s.Get(key)
		if err != nil {
			http.Error(writer, "hasn't succeeded to save the value", http.StatusInternalServerError)
		}
	}
}
