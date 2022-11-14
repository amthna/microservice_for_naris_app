package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// user struct (model)
type User struct {
	Firstname string `json:"first"`
	Email     string `json:"email"`
	Lastname  string `json:"last"`
	Status    string `json:"status"`
	Pay       int    `json:"pay"`
}

// Init users var as a slice User struct
var users []User

func initializeUsers(users *[]User, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &users)

}

// check for user
func checkUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	for _, item := range users {
		if (item.Email == user.Email) && (item.Firstname == user.Firstname) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("New user created:")
	users = append(users, user)

	file, _ := json.MarshalIndent(users, "", " ")

	_ = ioutil.WriteFile("employees.json", file, 0644)

	json.NewEncoder(w).Encode(user)

}

func main() {
	// fmt.Println("Hello wod")
	// init router
	r := mux.NewRouter()

	initializeUsers(&users, "employees.json")

	// mock datas @todo - implement DB
	// users = append(users, User{Firstname: "Frank", Lastname: "Stein", Email: "frank@hotmail.com"})
	// users = append(users, User{Firstname: "Count", Lastname: "Dracula", Email: "fangs@hotmail.com"})
	// users = append(users, User{Firstname: "Scary", Lastname: "Monster", Email: "monster@hotmail.com"})

	// route handlers/endpoints
	r.HandleFunc("/api/check", checkUser).Methods("POST")

	// run the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
