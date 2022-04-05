package server

import (
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
