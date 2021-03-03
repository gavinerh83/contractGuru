package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("./resource"))))
	http.HandleFunc("/", index)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8000", nil)
}

//main page
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//check if user is currently logged in
func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("myCookie")
	if err != nil {
		fmt.Println(err)
	}
	//send the cookie value for validation

}

func login(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		//redirect the user to the main page to sign in
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

}

func signup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	lastname := r.FormValue("lastname")
	firstname := r.FormValue("firstname")
	department := r.FormValue("department")
	//do some validation

	//create uuid for userid

	//save in server data structure

	//save in database
}
