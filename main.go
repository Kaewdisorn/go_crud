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
	Handlerequest()
	defer db.Close()
}

func Handlerequest() {

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	//http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	//templateData := map[string]interface{}{"Name": "Jon Snow"}
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	//tmpl.Execute(w, templateData)

}

func login(w http.ResponseWriter, r *http.Request) {

	var member Member
	if r.Method == "POST" {

		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username)
		fmt.Println(password)
		err := db.QueryRow("SELECT username, password FROM member where username = ?", username).
			Scan(&member.Username, &member.Password)
		if err != nil {
			//panic(err.Error())
			fmt.Println(err)
			templateData := map[string]interface{}{"Result": "Invalid Username"}
			tmpl.ExecuteTemplate(w, "index.gohtml", templateData)
		} else if username == member.Username && password == member.Password {
			fmt.Println("Login Success")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			templateData := map[string]interface{}{"Result": "Invalid Password"}
			tmpl.ExecuteTemplate(w, "index.gohtml", templateData)
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "Welcome to the HomePage!")
	tmpl.ExecuteTemplate(w, "register.gohtml", nil)
	//http.Redirect(w, r, "http://www.google.com", 301)
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		insertoDB(username, password, email)
	}
}

func insertoDB(username, password, email string) {

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(email)

}

/*func main() {

	fmt.Println("Server started on: http://localhost:8080")
	Handlerequest()
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

	//TEST Query Single Row
	var member Member
	// Execute the query
	err = db.QueryRow("SELECT * FROM member where id = ?", 2).Scan(&member.ID, &member.Username, &member.Password, &member.Email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("SINGLE ROW QUERY TEST")
	fmt.Println(member.ID, member.Username)
}*/
