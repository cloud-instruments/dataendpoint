package controllers

import (
	"cloudinstruments/web/dataproviders"
	"cloudinstruments/web/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var GetBatteryTestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("projectName")
		if projectName == "" {
			http.Error(w, "Invalid projectName", http.StatusInternalServerError)
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		provider.GetBatteryTest(projectName)
		fmt.Fprint(w, "GetBatteryTest executed")
	})

var PostBatteryTestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		d := json.NewDecoder(r.Body)
		var t models.Project
		err := d.Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print(t)
		provider := dataproviders.NewDynamoDBDataProvider()
		provider.PostBatteryTest(&t)
		fmt.Fprint(w, "PostBatteryTest executed")
	})

var DeleteBatteryTestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("projectName")
		if projectName == "" {
			http.Error(w, "Invalid projectName", http.StatusInternalServerError)
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		provider.DeleteBatteryTest(projectName)
		fmt.Fprint(w, "DeleteBatteryTest executed")
	})
