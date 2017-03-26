package controllers

import (
	"net/http"
)

var RootHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		resp :=
			`
		/GetToken - Generates a JWT token
		/PostProject - Saves a new battery project
		/DeleteProject - Deletes a battery project along with all associated data
		/GetProjectsByDeviceName - Retrieves a list of battery projects by device name
		/PostBatteryCycles - Save a battery cycle
		/GetBatteryCycles - Retrieves a list of battery cycles for ProjectName
		`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
	})
