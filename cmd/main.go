package main

import (
	"github.com/gorilla/mux"
	"github.com/sanjibgirics/bank-holiday/pkg"
	"log"
	"net/http"
)

func main() {
	// Fetch and process the bank holidays data.
	log.Println("Fetching and processing Data.")
	pkg.FetchAndProcessBankHolidayData()

	// Creates endpoints to serve the processed data
	log.Println("Serving the processed data through endpoints.")
	// Create router
	r := mux.NewRouter()

	// Register all the endpoints to the router
	pkg.RegisterRoutes(r)

	// Start listening and serving
	http.ListenAndServe(":8082", r)
}
