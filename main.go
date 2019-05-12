package main

import (
	"log"
	"net/http"

	"github.com/code-sleuth/yummy-recipes-go/models"

	"github.com/code-sleuth/yummy-recipes-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// run db adapter
	models.Adapter()

	mux := mux.NewRouter()
	mux.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	if err := http.ListenAndServe(":3333", mux); err != nil {
		log.Fatal(err)
	}

}
