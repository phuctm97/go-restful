package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List users\n")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Get user %s\n", vars["userId"])
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create user\n")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Update user %s\n", vars["userId"])
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Delete user %s\n", vars["userId"])
}

func main() {
	// Configure routes.
	router := mux.NewRouter()
	router.HandleFunc("/users/", listUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/", createUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{userId}/", getUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}/", updateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{userId}/", deleteUser).Methods(http.MethodDelete)

	// Start HTTP server.
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
