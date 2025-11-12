package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Plus(x, y int) int {
	return x + y
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("HELLO")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! You requested: %s", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
