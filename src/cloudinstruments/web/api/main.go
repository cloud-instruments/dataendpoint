package main

import (
	"cloudinstruments/web/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	resp :=
		`
	/PostBatteryTest - Save a new battery test cycle"
	/DeleteBatteryTest - Removes a new battery test by test id and test name
	/GetBatteryTest - Removes a new battery test by test id
	
	`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

func getBatteryTestHandler(w http.ResponseWriter, r *http.Request) {
	testId := r.URL.Query().Get("testId")
	if testId == "" {
		http.Error(w, "Invalid testId", http.StatusInternalServerError)
	}

	var provider BatteryDataProvider = NewDynamoDBDataProvider()
	provider.GetBatteryTest(testId)
	fmt.Fprint(w, "GetBatteryTest executed")
}

func postBatteryTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	d := json.NewDecoder(r.Body)
	var t models.BatteryCycle
	err := d.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print(t)
	var provider BatteryDataProvider = NewDynamoDBDataProvider()
	provider.PostBatteryTest(&t)
	fmt.Fprint(w, "PostBatteryTest executed")
}

func deleteBatteryTestHandler(w http.ResponseWriter, r *http.Request) {
	testId := r.URL.Query().Get("testId")
	if testId == "" {
		http.Error(w, "Invalid testId", http.StatusInternalServerError)
	}

	testName := r.URL.Query().Get("testName")
	if testName == "" {
		http.Error(w, "Invalid testName", http.StatusInternalServerError)
	}

	var provider BatteryDataProvider = NewDynamoDBDataProvider()
	provider.DeleteBatteryTest(testId, testName)
	fmt.Fprint(w, "DeleteBatteryTest executed")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/PostBatteryTest", postBatteryTestHandler)
	http.HandleFunc("/DeleteBatteryTest", deleteBatteryTestHandler)
	http.HandleFunc("/GetBatteryTest", getBatteryTestHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: ", err)
	}
}
