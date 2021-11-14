package main

import (
	"fmt"
	"net/http" // import package net-http
	"text/template"

	m "github.com/Kaewdisorn/module"
)

type Member struct {
	ID       int    //`json:"id"`
	Username string //`json:"username"`
	Password string //`json:"password"`
	Email    string //`json:"email"`
}

var tmpl = template.Must(template.ParseGlob("html/*")) // Declear variable for Html folder
var db = m.ConDB()

func main() {

	fmt.Println("Server started on: http://localhost:8080")
	//Handlerequest()
	//db := m.ConDB()
	//m.ConDB()
	defer db.Close()

	results, err := db.Query("SELECT * FROM member")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var member Member
		// for each row, scan the result into member object
		err = results.Scan(&member.ID, &member.Username, &member.Password, &member.Email)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the attribute
		fmt.Println(member.ID, member.Username, member.Password, member.Email)
	}
	defer results.Close()

	/*TEST Query Single Row*/
	var member Member
	// Execute the query
	err = db.QueryRow("SELECT * FROM member where id = ?", 2).Scan(&member.ID, &member.Username, &member.Password, &member.Email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("SINGLE ROW QUERY TEST")
	fmt.Println(member.ID, member.Username)
}

// Create method for client request
func homePage(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "Welcome to the HomePage!")
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func register(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "Welcome to the HomePage!")
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func Handlerequest() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/register", register)

	//http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)

}
