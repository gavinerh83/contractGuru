package main

import (
	"fmt"
	"go_Apprenticeship/secure"
	"go_Apprenticeship/user"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	tpl        *template.Template
	sessionMap = map[string]string{}    //JWT as the key, value as the username
	userMap    = map[string]user.Data{} //store user information, key is username, value is userinfo created in signup
)

func init() {
	//obtain data from database

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("./resource"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8000", nil)
}

//main page
func index(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//check if user is currently logged in
func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("myCookie")
	if err != nil {
		fmt.Println(err)
	}
	//send the cookie value for validation
	_, err = secure.ParseToken(cookie.Value)
	if err != nil {
		//Delete the session from map and delete cookie
		val := cookie.Value
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		delete(sessionMap, val)
		return false
	}
	return true
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		//redirect the user to the main page to sign in
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	//validate info
	var currentUser user.Data
	if _, ok := userMap[email]; !ok {
		tpl.ExecuteTemplate(w, "redirect.html", "User is not found")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(password))
	if err != nil {
		tpl.ExecuteTemplate(w, "redirect.html", "Wrong password entered")
		return
	}
	//generate sessionID
	id := uuid.NewV4()
	claim := &secure.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
		SessionID: id.String(),
	}
	signedToken, err := secure.GenerateJWT(claim)
	if err != nil {
		log.Println(err)
	}
	myCookie := &http.Cookie{
		Name:  "myCookie",
		Value: signedToken,
	}
	http.SetCookie(w, myCookie)
	//store signed JWT in sessionMap
	sessionMap[signedToken] = email
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	lastname := r.FormValue("lastname")
	firstname := r.FormValue("firstname")
	department := r.FormValue("department")
	//do some validation for unexpected characters

	bpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	//convert department string from form to integer
	departmentID := user.DepartmentMap[department]
	//create unique userid
	id := len(userMap)
	//save in server data structure
	userMap[email] = user.Data{Firstname: firstname, Password: string(bpassword), Lastname: lastname, Department: departmentID, Email: email, UserID: id}
	//save in database

}

func logout(w http.ResponseWriter, r *http.Request) {
	//link to logout button
}
