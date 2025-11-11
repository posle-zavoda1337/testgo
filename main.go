package main

import (
	"fmt"
	"log"
	"net/http"
)

func Plus(x int, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! You requested: %s", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
