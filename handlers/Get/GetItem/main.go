package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestBody struct {
	Serial     string `json:"serial"`
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse request body
	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		fmt.Println("Error parsing request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create URL
	apiURL := ""
	params := url.Values{}
	params.Set("userId", "3198")
	params.Set("serial", requestBody.Serial)
	

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
 	req.Header.Set("Authorization", "Bearer token")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
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
	http.HandleFunc("/api/viewAsset", handleAPIRequest)
	fmt.Println("Go server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}














