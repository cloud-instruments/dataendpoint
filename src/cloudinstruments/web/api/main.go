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
	http.HandleFunc("/PostBatteryTest", handlers.AuthHandler(controllers.PostBatteryTestHandler))
	http.HandleFunc("/DeleteBatteryTest", handlers.AuthHandler(controllers.DeleteBatteryTestHandler))
	http.HandleFunc("/GetBatteryTest", handlers.AuthHandler(controllers.GetBatteryTestHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
