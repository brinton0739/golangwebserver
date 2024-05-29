package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brinton0739/golangwebserver/internal/api"
	"github.com/gorilla/mux"
)

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/users", api.GetUsers)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"John Doe","email":"john@example.com"},{"id":2,"name":"Jane Doe","email":"jane@example.com"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
