package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
}

func main() {
	// Configure routes.
	router := mux.NewRouter()
	router.HandleFunc("/", handle)

	// Start HTTP server.
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
