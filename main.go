package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie/ressources"
)

func main() {
	http.HandleFunc("/", groupie.HandleHome)
	http.HandleFunc("/informations/{id}", groupie.HandleInfos)
	http.HandleFunc("/search", groupie.HandleSearch)
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Server failed to start: %v", err)
		return
	}
}
