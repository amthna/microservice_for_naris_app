package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// user struct (model)
type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// Init users var as a slice User struct
var users []User

// get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// get single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	for _, item := range users {
		if item.Email == params["email"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// create new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}

// update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.Email == params["email"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.Email = params["email"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}

	}
}

// delete user
func deleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.Email == params["email"] {
			users = append(users[:index], users[index+1:]...)
			break
		}

	}
}

func main() {
	// fmt.Println("Hello wod")
	// init router
	r := mux.NewRouter()

	// mock datas @todo - implement DB
	users = append(users, User{ID: "1", Firstname: "Frank", Lastname: "Stein", Email: "frank@hotmail.com"})
	users = append(users, User{ID: "2", Firstname: "Count", Lastname: "Dracula", Email: "fangs@hotmail.com"})
	users = append(users, User{ID: "3", Firstname: "Scary", Lastname: "Monster", Email: "monster@hotmail.com"})

	// route handlers/endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{email}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{email}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{email}", deleteUsers).Methods("DELETE")

	// run the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
