package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"banco/server"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods("POST")
	router.HandleFunc("/users", server.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", server.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
		
	
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	
}
