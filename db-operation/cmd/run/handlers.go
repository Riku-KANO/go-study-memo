package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"github.com/Riku-KANO/db-operation/models"
)

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running application!")
}

func (a *App) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}

	res, err := a.models.Users.GetAll()
	if err != nil {
		log.Fatal(err)
		io.WriteString(w, "error occured")
		return
	}
	
	for _, user := range res {
		io.WriteString(w, user.Name)
	}
}

func (a *App) createUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}

	var userRequest models.UserRequest
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user := models.User{
		Name: userRequest.Name,
	}

	err = a.models.Users.Create(&user)
	if err != nil {
		log.Fatal(err)
		io.WriteString(w, "errors occured")
		return
	}

	log.Println("new user created")
}