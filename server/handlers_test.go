package server

import (
	"github.com/Alxus228/Key-Value/storage"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	testStorage := storage.New()
	testStorage.Put(1, 0)
	testStorage.Put("2", "a")

	r, err := http.NewRequest("GET", "/api/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	vars := map[string]string{
		"key": "1",
	}
	r = mux.SetURLVars(r, vars)
	getHandler(testStorage)(w, r)

	if w.Code != http.StatusOK {
		bodyBytes, _ := io.ReadAll(w.Body)
		t.Error(testStorage)
		t.Error(string(bodyBytes))
		t.Errorf("Expected status code %d, but got: %d", http.StatusOK, w.Code)
	}
}
