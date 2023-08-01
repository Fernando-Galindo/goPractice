package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	// Define the API endpoint handler function
	http.HandleFunc("/healthy", checkHealth)

	// Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	// Simulating a database query
	msg := map[string]string{"status": "healthy", "date": time.Now().GoString()}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Convert the users slice to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonData)
}
