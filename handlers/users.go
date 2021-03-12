package handlers

import (
	"fmt"
	"go_Apprenticeship/models"
	"go_Apprenticeship/secure"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

//Index Serves the prelogged in page
func Index(w http.ResponseWriter, r *http.Request) {
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
		cookie = &http.Cookie{
			Name:  "myCookie",
			Value: "",
		}
	}
	http.SetCookie(w, cookie)
	//send the cookie value for validation
	_, err = secure.ParseToken(cookie.Value)
	if err != nil {
		//Delete the session from map
		delete(models.SessionMap, cookie.Value)
		return false
	}
	return true
}

//LoginUser receives post form input
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		//redirect the user to the main page to sign in
		http.Redirect(w, r, "/api/v1/users", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		//validate info
		currentUser := models.UserMap[email]
		if _, ok := models.UserMap[email]; !ok {
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
		models.SessionMap[signedToken] = email
		http.Redirect(w, r, "/api/v1/users/main", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/api/v1/users", http.StatusSeeOther)
}

//SignupUser receives post form values
func SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("signupEmail")
		password := r.FormValue("signupPassword")
		lastname := r.FormValue("lastname")
		firstname := r.FormValue("firstname")
		department := r.FormValue("department")
		//do some validation for unexpected characters

		bpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			fmt.Println(err)
		}
		//convert department string from form to integer
		departmentID := models.DepartmentMap[department]
		//create unique userid
		id := len(models.UserMap)
		//save in server data structure
		models.UserMap[email] = models.Data{Firstname: firstname, Password: string(bpassword), Lastname: lastname, Department: departmentID, Email: email, UserID: id}
		//save in database

		err = tpl.ExecuteTemplate(w, "redirect.html", "Please sign in again")
	}
	http.Redirect(w, r, "/api/v1/users", http.StatusSeeOther)
}

//LogoutUser logs the user out
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	//link to logout button

}

//Home displays the main page
func Home(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		tpl.ExecuteTemplate(w, "redirect.html", "Not logged in yet")
		return
	}
	tpl.ExecuteTemplate(w, "home.html", nil)
}
