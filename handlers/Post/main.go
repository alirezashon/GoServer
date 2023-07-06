package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateAssetRequestBody struct {
	UserID         string `json:"userId"`
	AssetCode      string `json:"assetcode"`
	Status         string `json:"status"`
	Location       string `json:"location"`
	AllottedTo     string `json:"allottedto"`
	IsCustomUpdate bool   `json:"isCustomUpdate"`
	Remark         string `json:"remark"`
}

type LinkAssetRequestBody struct {
	UserID    string `json:"userId"`
	AssetList []struct {
		AssetCode      string `json:"assetcode"`
		LinkAssetCode  string `json:"linkassetcode"`
	} `json:"assetList"`
}

type AssetInstallationRequestBody struct {
	UserID           string `json:"userId"`
	AssetCode        string `json:"assetcode"`
	Status           string `json:"status"`
	Location         string `json:"location"`
	CustomerDetails  string `json:"CustomerDetails"`
	ChildAssetCode   string `json:"ChildAssetcode"`
	CustomerName     string `json:"CustomerName"`
	AgentCode        string `json:"AgentCode"`
	IsRemoveLink     string `json:"IsRemoveLink"`
	AllottedTo       string `json:"AllottedTo"`
	SettlementType   string `json:"SettlementType"`
}

func makeRequest(url string, requestBody interface{}) ([]byte, error) {
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
 	req.Header.Set("Authorization", "Bearer token")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
func updateAsset(w http.ResponseWriter, r *http.Request) {
	// Log the request
	log.Println("Received updateAsset request")

	// Decode the request body
	var requestBody UpdateAssetRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add additional fields to the request body
	requestBody.UserID = "3198"
	requestBody.Location = "example-location"
	requestBody.AllottedTo = "customer@test.ir"
	requestBody.IsCustomUpdate = true

	apiURL := "http://10.104.26.112:81/api/updateAssetStatus"

	// Make the request
	response, err := makeRequest(apiURL, requestBody)
	if err != nil {
		log.Println("Error making request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log the response
	log.Println("Received response from updateAsset API")

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// func updateAsset(w http.ResponseWriter, r *http.Request) {
// 	var requestBody UpdateAssetRequestBody
// 	err := json.NewDecoder(r.Body).Decode(&requestBody)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	requestBody.UserID = "3198"
// 	requestBody.Location = "example-location"
// 	requestBody.AllottedTo = "customer@test.ir"
// 	requestBody.IsCustomUpdate = true

// 	apiURL := "http://10.104.26.112:81/api/updateAsset"

// 	response, err := makeRequest(apiURL, requestBody)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(response)
// }

func linkAsset(w http.ResponseWriter, r *http.Request) {
	var requestBody LinkAssetRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	requestBody.UserID = "3198"

	apiURL := "http://10.104.26.112:81/api/linkAsset"

	response, err := makeRequest(apiURL, requestBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func assetInstallation(w http.ResponseWriter, r *http.Request) {
	var requestBody AssetInstallationRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	requestBody.UserID = "3198"
	requestBody.Location = "example-location"

	apiURL := "http://10.104.26.112:81/api/AssetInstallation"

	response, err := makeRequest(apiURL, requestBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/api/updateAsset", updateAsset)
	http.HandleFunc("/api/linkAsset", linkAsset)
	http.HandleFunc("/api/assetInstallation", assetInstallation)

	fmt.Println("Go server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
