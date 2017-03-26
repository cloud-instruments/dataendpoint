package main

import (
	"cloudinstruments/web/controllers"
	"cloudinstruments/web/handlers"
	"fmt"
	"net/http"
	"os"
)

var (
	authSecret, jwtSecret []byte
)

func main() {
	if os.Getenv("ENV") == "dev" {
		http.HandleFunc("/", controllers.RootHandler)
		http.HandleFunc("/GetToken", controllers.GetTokenHandler)
		//projects
		http.HandleFunc("/PostProject", controllers.PostProjectHandler)
		http.HandleFunc("/DeleteProject", controllers.DeleteProjectHandler)
		http.HandleFunc("/GetProjectsByDeviceName", controllers.GetProjectsHandler)
		//cycles
		http.HandleFunc("/PostProjectCycles", controllers.PostBatteryCycleHandler)
		http.HandleFunc("/GetProjectCycles", controllers.GetProjectCyclesHandler)
		// todo delete battery cycles
		http.HandleFunc("/DeleteProjectyCycles", controllers.DeleteProjectCyclesHandler)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("ListenAndServe: ", err)
		}
	} else {
		http.HandleFunc("/", handlers.AuthHandler(controllers.RootHandler))
		http.HandleFunc("/GetToken", controllers.GetTokenHandler)
		//projects
		http.HandleFunc("/PostProject", handlers.AuthHandler(controllers.PostProjectHandler))
		http.HandleFunc("/DeleteProject", handlers.AuthHandler(controllers.DeleteProjectHandler))
		http.HandleFunc("/GetProjectsByDeviceName", handlers.AuthHandler(controllers.GetProjectsHandler))
		//cycles
		http.HandleFunc("/PostProjectCycles", handlers.AuthHandler(controllers.PostBatteryCycleHandler))
		http.HandleFunc("/GetProjectCycles", handlers.AuthHandler(controllers.GetProjectCyclesHandler))
		// todo delete battery cycles
		http.HandleFunc("/GetProjectCycles", handlers.AuthHandler(controllers.DeleteProjectCyclesHandler))
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("ListenAndServe: ", err)
		}
	}
}
