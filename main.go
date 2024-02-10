package main

import (
	"log"
	"net/http"
)

//Setup vars and slices

//Setup structs

//Utility Functions

// Core Functions
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page!"))
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the about Page."))
}

// Main Function
func main() {
	//Register routes Handlerfunctions in main function
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)

	//Start server
	println("Server is running on port 8080...")
	//error check
	//Simple log out
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start doto error: %v", err)
	}
}
