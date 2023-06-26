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
 	req.Header.Set("Authorization", "Bearer  lHvpvXqxo6g38q2Uzi8KrTPNar_oW129dM3w2daLKOhKkrPJF1sNN4WlaaJe2qX86fJx0ga2Q8gsH7hgBbdtQcKrb3icoJODW6o2pZOYa5IQMP_NpZEAc2JIGA-WSagESVmPOww3le4Ag99c5IuSWEtoF3McGTHw2wgTnAyoeYcuEcSn-Wmz57dydBSsg5bY44A2YZNZrwbZFaCFzvqY3yZdDjinqdMqp94nDenV1X1K-oG1MRuegGrx33jmgzZ3B5QtOPHFzXSmu7KnZZydJsYV_LyWI-aLNvd51EUYg_z3XtgELRAPC0Odz56A-QPkE7wuKgtn9mDtG32L5RjS7Q_MDcjgnSXxVPZEhsLcmxVgGDPO4oO7-n2nMDL7gb_N7HYPFqRsUQsX9_JPYqDZkzE_ImIWOoYw97JBsiTuXzygBFQeQbHkh75z_CZUY0CBTor83ooTT_uIPwy9kmy0t9maNuxWN7OSW95wibv_rJMTTxiVgdCQnddYMLNMyxWFVodInetbzQQ-6ohEWUifvTKWq2hKrWCwc-jFqn3lIuFnJDQhzvIveif11lrEKj-yQYxk8eG8vUwOJbjDuz-JdAl68D-AKSsbKf1QqOTYwlv-qboKL3wjaPuFRnvtYUnkDfjptvQlTCszgccXvBL2cH9TpTAnTrRYChWVqOy8byTUjPoNvbPvQ9g4TSnuvhd6mnKZOpIxXhBQaqiQHhLnQUdhO-DrzX5ONzu1B_BGMGao81O_8hy6WrYx9Pz18tfN2qdSI55B121uUs7fbEyO8zxGFGjKjLCsS8qV25_Zgs-CKXv8lE6njVXKp84d-H-2ZdlIjENbp83OmIej0U3T9sngCRTBTSu7kopQcpDk5UGP48goqwrrJ41Y3nUdVUW9UCRFyqZV8pX3XFnn-Uh74juf75Wvk-8Vqpl2EQbKCd3Ka7my56ALsiB3K1cNWVUfCHmluHwc0NzwQSzY4jdEfcZlj9KVf3rcmZ94bWQnXng1OYI7QCJnVyGjDcTnaHe58-WJ3srCJg9t6dxRly-Iy-6obFSUfldX3QOEmYO4pGwLdMCvrcKhhmFMySNdT6MDPrrAsQ")

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
