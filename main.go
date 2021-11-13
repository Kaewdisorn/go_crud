package main

import (
	"fmt"
	"net/http" // import package net-http
)

func main() {

	handlerequest()
}

// Create method for client request

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to the HomePage!")
}

func handlerequest() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
