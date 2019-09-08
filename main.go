package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
}

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
