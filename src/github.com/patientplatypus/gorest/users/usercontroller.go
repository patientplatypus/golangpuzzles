package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	//super duper dangerous - Axe in production!
	// log.Print("Waiting for create table")
	// finished := make(chan bool)
	// go Create_Table(finished, "classes")
	// <-finished
	// log.Print("Continuing with the rest of the program")

	DeleteAll()
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished, "classes")
	<-finished
	log.Print("Continuing with the rest of the program")

	if username != "" && password != "" {
		user := User{Username: username, Password: password}
		user.Create()
	} else {
		fmt.Fprintln(w, "error username or password is blank!")
	}
}

type User struct {
	Username string
	Password string
	Id       int
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var loginjson User
	err := decoder.Decode(&loginjson)

	if err != nil {
		panic(err)
	}

	username := loginjson.Username
	password := loginjson.Password

	log.Print("username: ", username)
	log.Print("password: ", password)
	if username != "" && password != "" {
		loginjson.Login(w, r)
	} else {
		fmt.Fprintln(w, "error username or password is blank!")
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	Logout(w, r)
}

func QueryAllUsers(w http.ResponseWriter, r *http.Request) {
	QueryAll()
}
