package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Printf("Request: %#v", r.Form)
	})
	http.ListenAndServe(":8080", nil)
}
