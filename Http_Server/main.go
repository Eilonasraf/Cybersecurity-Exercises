package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for processing form submissions via POST requests
// Extracts the 'name' and 'address' form values and prints them.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Respond to the client
	fmt.Fprintf(w, "POST request successful\n")

	// Extract form values for "name" and "address"
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Print the form values back to the client
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Handler for the "/hello" route
// Responds with a "hello!" message if the request is valid.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is exactly "/hello", otherwise return a 404 error.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Only allow GET requests, return a 404 error for other methods.
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Respond with a hello message
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Serve static files from the "static" directory at the root path "/".
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Handle form submissions at the "/form" route
	http.HandleFunc("/form", formHandler)

	// Handle GET requests at the "/hello" route
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // Log any error that occurs while starting the server
	}
}