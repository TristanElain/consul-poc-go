package main

import (
	"consul-poc/controller"
	"consul-poc/utils"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting app...")
	baseController := controller.BaseController{}

	http.HandleFunc("/", baseController.HandleRoot)
	http.HandleFunc("/map", baseController.HandleMapProps)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	fmt.Println("listening on port 4000")
	err := http.ListenAndServe("localhost:4000", nil)
	utils.FatalError(err)
}
