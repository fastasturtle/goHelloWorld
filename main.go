package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Starting...")
	http.ListenAndServe(":8080", nil)
}
