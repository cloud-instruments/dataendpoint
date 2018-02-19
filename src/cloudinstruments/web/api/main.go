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
