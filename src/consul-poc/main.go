package main

import (
	"consul-poc/controller"
	"consul-poc/manager"
	"consul-poc/utils"
	"fmt"
	"net/http"
)

// http://varunksaini.com/consul-service-discovery-golang/
func main() {
	fmt.Println("Starting app...")
	baseController := controller.BaseController{}
	serviceManager := manager.GetConsulServiceManager()

	// Endpoints
	http.HandleFunc("/", baseController.HandleForm)
	http.HandleFunc("/configuration", baseController.HandleProps)
	http.HandleFunc("/configuration/map", baseController.HandleMapProps)
	http.HandleFunc("/health", baseController.HandleHealthCheck)

	// Registering to consul
	serviceManager.RegisterService("webservice-go", 4000)
	// serviceManager.DeregisterService("webservice-go-1522936919")

	fmt.Println("listening on port 4000")
	err := http.ListenAndServe("localhost:4000", nil)
	utils.FatalError(err)
}
