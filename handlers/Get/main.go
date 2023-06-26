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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer  lHvpvXqxo6g38q2Uzi8KrTPNar_oW129dM3w2daLKOhKkrPJF1sNN4WlaaJe2qX86fJx0ga2Q8gsH7hgBbdtQcKrb3icoJODW6o2pZOYa5IQMP_NpZEAc2JIGA-WSagESVmPOww3le4Ag99c5IuSWEtoF3McGTHw2wgTnAyoeYcuEcSn-Wmz57dydBSsg5bY44A2YZNZrwbZFaCFzvqY3yZdDjinqdMqp94nDenV1X1K-oG1MRuegGrx33jmgzZ3B5QtOPHFzXSmu7KnZZydJsYV_LyWI-aLNvd51EUYg_z3XtgELRAPC0Odz56A-QPkE7wuKgtn9mDtG32L5RjS7Q_MDcjgnSXxVPZEhsLcmxVgGDPO4oO7-n2nMDL7gb_N7HYPFqRsUQsX9_JPYqDZkzE_ImIWOoYw97JBsiTuXzygBFQeQbHkh75z_CZUY0CBTor83ooTT_uIPwy9kmy0t9maNuxWN7OSW95wibv_rJMTTxiVgdCQnddYMLNMyxWFVodInetbzQQ-6ohEWUifvTKWq2hKrWCwc-jFqn3lIuFnJDQhzvIveif11lrEKj-yQYxk8eG8vUwOJbjDuz-JdAl68D-AKSsbKf1QqOTYwlv-qboKL3wjaPuFRnvtYUnkDfjptvQlTCszgccXvBL2cH9TpTAnTrRYChWVqOy8byTUjPoNvbPvQ9g4TSnuvhd6mnKZOpIxXhBQaqiQHhLnQUdhO-DrzX5ONzu1B_BGMGao81O_8hy6WrYx9Pz18tfN2qdSI55B121uUs7fbEyO8zxGFGjKjLCsS8qV25_Zgs-CKXv8lE6njVXKp84d-H-2ZdlIjENbp83OmIej0U3T9sngCRTBTSu7kopQcpDk5UGP48goqwrrJ41Y3nUdVUW9UCRFyqZV8pX3XFnn-Uh74juf75Wvk-8Vqpl2EQbKCd3Ka7my56ALsiB3K1cNWVUfCHmluHwc0NzwQSzY4jdEfcZlj9KVf3rcmZ94bWQnXng1OYI7QCJnVyGjDcTnaHe58-WJ3srCJg9t6dxRly-Iy-6obFSUfldX3QOEmYO4pGwLdMCvrcKhhmFMySNdT6MDPrrAsQ")

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
