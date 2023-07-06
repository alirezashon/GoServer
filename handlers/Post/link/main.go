
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func linkAsset(w http.ResponseWriter, r *http.Request) {
	// Log the request
	log.Println("Received AssetInstallation request")
	
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	// Create a new request to the asset API
	req, err := http.NewRequest("POST", "http://10.104.26.112:81/api/linkAsset", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	// Perform the HTTP request to the asset API
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log the response
	log.Println("Responding to AssetInstallation request")

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func main() {
	http.HandleFunc("/api/linkAsset", linkAsset )

	fmt.Println("Go server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
