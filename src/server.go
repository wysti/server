package main

import (
	"fmt"
	"net/http"
	"rest"
)

func initializeHandlers() []rest.RestHandler {
	restHandlers := make([]rest.RestHandler, 0)
	restHandlers = append(restHandlers, rest.NewBasicSearchHandler())
	return restHandlers
}

func acceptRequests(restHandlers []rest.RestHandler) {
	mux := http.NewServeMux()
	for _, rh := range restHandlers {
		mux.HandleFunc(rh.Context(), rh.Handler)
	}
	fmt.Println("Listening...")
	http.ListenAndServe(":9090", mux)
}

func main() {
	//initializeDB()
	//initializeLogging()
	rest.InitializeErrors()
	restHandlers := initializeHandlers()
	acceptRequests(restHandlers)
}
