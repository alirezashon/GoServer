package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
//13008248
func main() {
	agentID := "10000000"
	accessories := "MMD"

	// Create URL with query parameters
	apiURL := "http://10.104.26.112:81/api/viewAsset"
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
		return
	}


	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse JSON response
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print response
	fmt.Println(data)
}
