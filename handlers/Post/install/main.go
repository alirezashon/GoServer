

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func AssetInstallation(w http.ResponseWriter, r *http.Request) {
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
	req, err := http.NewRequest("POST", "http://10.104.26.112:81/api/AssetInstallation", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer  lHvpvXqxo6g38q2Uzi8KrTPNar_oW129dM3w2daLKOhKkrPJF1sNN4WlaaJe2qX86fJx0ga2Q8gsH7hgBbdtQcKrb3icoJODW6o2pZOYa5IQMP_NpZEAc2JIGA-WSagESVmPOww3le4Ag99c5IuSWEtoF3McGTHw2wgTnAyoeYcuEcSn-Wmz57dydBSsg5bY44A2YZNZrwbZFaCFzvqY3yZdDjinqdMqp94nDenV1X1K-oG1MRuegGrx33jmgzZ3B5QtOPHFzXSmu7KnZZydJsYV_LyWI-aLNvd51EUYg_z3XtgELRAPC0Odz56A-QPkE7wuKgtn9mDtG32L5RjS7Q_MDcjgnSXxVPZEhsLcmxVgGDPO4oO7-n2nMDL7gb_N7HYPFqRsUQsX9_JPYqDZkzE_ImIWOoYw97JBsiTuXzygBFQeQbHkh75z_CZUY0CBTor83ooTT_uIPwy9kmy0t9maNuxWN7OSW95wibv_rJMTTxiVgdCQnddYMLNMyxWFVodInetbzQQ-6ohEWUifvTKWq2hKrWCwc-jFqn3lIuFnJDQhzvIveif11lrEKj-yQYxk8eG8vUwOJbjDuz-JdAl68D-AKSsbKf1QqOTYwlv-qboKL3wjaPuFRnvtYUnkDfjptvQlTCszgccXvBL2cH9TpTAnTrRYChWVqOy8byTUjPoNvbPvQ9g4TSnuvhd6mnKZOpIxXhBQaqiQHhLnQUdhO-DrzX5ONzu1B_BGMGao81O_8hy6WrYx9Pz18tfN2qdSI55B121uUs7fbEyO8zxGFGjKjLCsS8qV25_Zgs-CKXv8lE6njVXKp84d-H-2ZdlIjENbp83OmIej0U3T9sngCRTBTSu7kopQcpDk5UGP48goqwrrJ41Y3nUdVUW9UCRFyqZV8pX3XFnn-Uh74juf75Wvk-8Vqpl2EQbKCd3Ka7my56ALsiB3K1cNWVUfCHmluHwc0NzwQSzY4jdEfcZlj9KVf3rcmZ94bWQnXng1OYI7QCJnVyGjDcTnaHe58-WJ3srCJg9t6dxRly-Iy-6obFSUfldX3QOEmYO4pGwLdMCvrcKhhmFMySNdT6MDPrrAsQ")

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
	http.HandleFunc("/api/AssetInstallation", AssetInstallation)

	fmt.Println("Go server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
