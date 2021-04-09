package main

import (
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
	})

	http.ListenAndServe(":"+port, nil)
}
