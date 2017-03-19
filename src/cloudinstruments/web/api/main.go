package main

import (
	"cloudinstruments/web/controllers"
	"cloudinstruments/web/handlers"
	"fmt"
	"net/http"
)

var (
	authSecret, jwtSecret []byte
)

func main() {
	http.HandleFunc("/", handlers.AuthHandler(controllers.RootHandler))
	http.HandleFunc("/GetToken", controllers.GetTokenHandler)
	http.HandleFunc("/PostProject", handlers.AuthHandler(controllers.PostProjectHandler))
	http.HandleFunc("/DeleteProject", handlers.AuthHandler(controllers.DeleteBatteryTestHandler))
	http.HandleFunc("/GetBatteryProjects", handlers.AuthHandler(controllers.GetBatteryTestHandler))
	http.HandleFunc("/PostBatteryCycle", handlers.AuthHandler(controllers.PostBatteryCycleHandler))
	http.HandleFunc("/GetBatteryCycles", handlers.AuthHandler(controllers.PostBatteryCycleHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
