package main

import (
	"fmt"
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
		fmt.Printf("Error: Unable to start server on port 8080: %v\n", err)
		return
	}
}
