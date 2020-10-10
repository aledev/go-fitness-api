package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func initUsersRepo() {
	for i := 1; i <= 10000; i++ {
		Users = append(Users,
			User{Id: i,
				Name:     fmt.Sprintf("User %d", i),
				Username: fmt.Sprintf("username%d", i),
				Password: "123456"})
	}

	fmt.Println("Users were created!")
}

func userLogin(username string, password string) *User {
	for _, user := range Users {
		if user.Username == username &&
			user.Password == password {
			return &user
		}
	}
	return nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1000 * time.Millisecond)
	user := userLogin("username1", "123456")

	if user != nil {
		json.NewEncoder(w).Encode(&user)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", handleLogin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	var syncOnce sync.Once

	syncOnce.Do(initUsersRepo)

	handleRequests()
}
