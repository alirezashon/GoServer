package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBody struct {
	UserID    string `json:"userId"`
	AssetCode string `json:"assetcode"`
	Status    string `json:"status"`
	Location  string `json:"Location"`
	// Include other fields as needed
}

func makeRequest(url string, requestBody RequestBody) ([]byte, error) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func updateAsset(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Make API call to updateAsset
	url := "https://example.com/updateAsset" // Replace with the actual API endpoint
	responseBody, err := makeRequest(url, requestBody)
	if err != nil {
		log.Println("Error making API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func linkAsset(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Make API call to linkAsset
	url := "https://example.com/linkAsset" // Replace with the actual API endpoint
	responseBody, err := makeRequest(url, requestBody)
	if err != nil {
		log.Println("Error making API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func updateStatus(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Make API call to updateStatus
	url := "https://example.com/updateStatus" // Replace with the actual API endpoint
	responseBody, err := makeRequest(url, requestBody)
	if err != nil {
		log.Println("Error making API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func main() {
	http.HandleFunc("/updateAsset", updateAsset)
	http.HandleFunc("/linkAsset", linkAsset)
	http.HandleFunc("/updateStatus", updateStatus)

	fmt.Println("Go server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
