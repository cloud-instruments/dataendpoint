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

var PostBatteryCycleHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		d := json.NewDecoder(r.Body)
		var cycle models.BatteryCycle
		err := d.Decode(&cycle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print(cycle)
		provider := dataproviders.NewDynamoDBDataProvider()
		if _, errPosting := provider.PostBatteryCycle(&cycle); errPosting != nil {
			http.Error(w, "Invalid request method", http.StatusInternalServerError)
		}

		fmt.Fprint(w, "PostBatteryCycle executed")
	})

var PostProjectHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		d := json.NewDecoder(r.Body)
		var project models.Project
		err := d.Decode(&project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print(project)
		provider := dataproviders.NewDynamoDBDataProvider()
		if _, errPosting := provider.PostProject(&project); errPosting != nil {
			http.Error(w, "Invalid request method", http.StatusInternalServerError)
		}

		fmt.Fprint(w, "PostProject executed")
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
