package module

import (
	"fmt"
	"net/http" // import package net-http
)

// Create method for client request
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to the HomePage!")
}

func Handlerequest() {

	//http.HandleFunc("/", homePage)
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)

}
