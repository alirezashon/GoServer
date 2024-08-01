package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	agentID := "10000000"
	accessories := "mmd"

	// Create URL with query parameters
	apiURL := ""
	params := url.Values{}
	params.Set("userId", "3198")
	params.Set("status", "Contractor---Round up")
	params.Set("agentCode", agentID)
	params.Set("Category", accessories)

	urlWithParams := apiURL + "?" + params.Encode()

	// Create HTTP client
	client := http.Client{}

	// Create GET request
	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token" )
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse JSON response
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/api/viewAsset", handleAPIRequest)
	fmt.Println("Go server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
