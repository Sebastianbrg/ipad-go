package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type City struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		cities := []City{
			{Name: "Paris", Country: "France"},
			{Name: "New York", Country: "United States"},
			{Name: "London", Country: "United Kingdom"},
			{Name: "Tokyo", Country: "Japan"},
			{Name: "Sydney", Country: "Australia"},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cities)
	})

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
