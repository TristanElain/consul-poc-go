package main

import (
	"consul-poc/controller"
	"consul-poc/registering"
	"consul-poc/utils"
	"fmt"
	"net/http"
)

// http://varunksaini.com/consul-service-discovery-golang/
func main() {
	fmt.Println("Starting app...")
	baseController := controller.BaseController{}
	serviceManager, err := registering.NewConsulServiceManager()
	utils.FatalError(err)

	// Endpoints
	http.HandleFunc("/", baseController.HandleRoot)
	http.HandleFunc("/map", baseController.HandleMapProps)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	// Registering to consul
	serviceManager.RegisterService("webservice-go", 4000)
	// serviceManager.DeregisterService("webservice-go-1522936919")

	fmt.Println("listening on port 4000")
	err = http.ListenAndServe("localhost:4000", nil)
	utils.FatalError(err)
}
