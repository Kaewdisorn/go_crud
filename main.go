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
var username string

func main() {

	fmt.Println("Server started on: http://localhost:8080")
	Handlerequest()
	defer db.Close()
}

func Handlerequest() {

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/insdb", insdb)
	http.HandleFunc("/member", show)
	http.HandleFunc("/deleteUsers", deleteUsers)
	//http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	var member Member
	var templateData = map[string]interface{}{}
	if r.Method == "POST" {

		username = r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username)
		fmt.Println(password)

		err := db.QueryRow("SELECT username, password FROM member where username = ?", username).
			Scan(&member.Username, &member.Password)
		if err != nil {
			//panic(err.Error())
			fmt.Println(err)
			templateData = map[string]interface{}{"Result": "Invalid Username"}
		} else if username == member.Username && password == member.Password {
			fmt.Println("Login Success")
			//templateData = map[string]interface{}{"Result": "Login Success"}
			//templateData = map[string]interface{}{"Uname": username}
			http.Redirect(w, r, "/member", http.StatusSeeOther)
			//http.Redirect(w, r, "/", http.StatusSeeOther)

		} else {
			templateData = map[string]interface{}{"Result": "Invalid Password"}
		}
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", templateData)
}

func register(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func insdb(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		ins, err := db.Prepare("INSERT INTO member(username,password,email) VALUES(?,?,?)")
		if err != err {
			panic(err.Error())
		} else {
			ins.Exec(username, password, email)
			fmt.Println("Insert success")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func show(w http.ResponseWriter, r *http.Request) {

	//templateData := map[string]interface{}{"Uname": username}
	selDB, err := db.Query("SELECT id,username,email FROM member")
	if err != nil {
		panic(err.Error())
	} else {
		users := make([]Member, 0)
		for selDB.Next() {
			usr := Member{}
			selDB.Scan(&usr.ID, &usr.Username, &usr.Email)
			users = append(users, usr)
		}
		fmt.Println(users)
		//tmpl.ExecuteTemplate(w, "member.gohtml",users)
		tmpl.ExecuteTemplate(w, "member.gohtml", map[string]interface{}{"Uname": username, "Ww": users})
	}
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Please send ID", http.StatusBadRequest)
	}
	_, er := db.Exec("DELETE FROM member WHERE id = ?", id)
	if er != nil {
		panic(er.Error())
	}
	http.Redirect(w, r, "/member", http.StatusSeeOther)
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
}*/
