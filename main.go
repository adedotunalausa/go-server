package main

import (
	"fmt"
	"log"
	"net/http"
)

var notFoundError = "404 not found!"
var methodNotAllowedError = "Method not allowed"

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHello(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(responseWriter, notFoundError, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(responseWriter, methodNotAllowedError, http.StatusForbidden)
		return
	}
	_, err := fmt.Fprintf(responseWriter, "Hello from Marxes")
	if err != nil {
		log.Fatalf("Error while greeting: %v", err)
		return
	}
}

func handleForm(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/form" {
		http.Error(responseWriter, notFoundError, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		http.Error(responseWriter, methodNotAllowedError, http.StatusForbidden)
		return
	}
	if err := request.ParseForm(); err != nil {
		_, err = fmt.Fprintf(responseWriter, "Error while parsing form: %s", err)
		if err != nil {
			return
		}
		return
	}
	_, err := fmt.Fprintf(responseWriter, "Form submitted successfully \n")
	if err != nil {
		log.Fatalf("Error while submitting form: %v", err)
		return
	}
	name := request.FormValue("name")
	address := request.FormValue("address")
	_, err = fmt.Fprintf(responseWriter, "Name = %s\n", name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(responseWriter, "Address = %s\n", address)
	if err != nil {
		return
	}
}
