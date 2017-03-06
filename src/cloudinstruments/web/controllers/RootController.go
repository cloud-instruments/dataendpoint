package controllers

import (
	"net/http"
)

var RootHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		resp :=
			`
		/GetToken - Generates a JWT token
		/PostBatteryTest - Save a new battery test
		/DeleteBatteryTest - Removes a new battery test by ProjectName
		/GetBatteryTest - Removes a new battery project by ProjectName
		
		`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
	})
