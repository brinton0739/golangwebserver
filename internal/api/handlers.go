package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

)

type User struct {
	ID	string `json:"id"`
	Name	string `json:"name`
	Email string `json:"email"`

}

var users = []User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
}

func RegisterRoutes(r *mux.Router) {

	r.Handlefunc("/users", GetUsers).Methods("GET")
	r.Handlefunc("/users/{id}", GetUser).Methods("GET")
	r.Handlefunc("/users", CreateUser).Methods("POST")
	r.Handlefunc("/users/{id}", UpdateUser).Methods("PUT")
	r.Handlefunc("/users/{id}", DeleteUser).Methods("DELETE")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "applcation/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)

		}
	}
}

func DeleteUser(w http.ResponseWriter r * http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}