Copyright (c) <2018> <University of Washington>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


package controllers

import (
	"cloudinstruments/web/dataproviders"
	"cloudinstruments/web/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var GetProjectsHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		deviceName := r.URL.Query().Get("deviceName")
		if deviceName == "" {
			http.Error(w, "Invalid deviceName", http.StatusInternalServerError)
			return
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		if projects, err := provider.GetProjectsByDeviceName(deviceName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			if resp, errMarshaling := json.Marshal(projects); errMarshaling != nil {
				http.Error(w, errMarshaling.Error(), http.StatusInternalServerError)
				return
			} else {
				w.Write(resp)
			}
		}
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
			http.Error(w, errPosting.Error(), http.StatusInternalServerError)
		}
	})

var DeleteProjectHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("projectName")
		if projectName == "" {
			http.Error(w, "Invalid projectName", http.StatusInternalServerError)
			return
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		if _, err := provider.DeleteProject(projectName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
			http.Error(w, errPosting.Error(), http.StatusInternalServerError)
		}
	})

var GetProjectCyclesHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("projectName")
		if projectName == "" {
			http.Error(w, "Invalid projectName", http.StatusInternalServerError)
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		if cycles, err := provider.GetProjectCyclesByProjectName(projectName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			if resp, errMarshaling := json.Marshal(cycles); errMarshaling != nil {
				http.Error(w, errMarshaling.Error(), http.StatusInternalServerError)
				return
			} else {
				w.Write(resp)
			}
		}
	})

var DeleteProjectCyclesHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("projectName")
		if projectName == "" {
			http.Error(w, "Invalid projectName", http.StatusInternalServerError)
		}

		provider := dataproviders.NewDynamoDBDataProvider()
		provider.DeleteProjectCycles(projectName)
	})
