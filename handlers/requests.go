package handlers

import "net/http"

//Requests routes the user to different url based on method type
func Requests(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/api/v1/getrequests", http.StatusSeeOther)
	} else if r.Method == http.MethodPost {
		http.Redirect(w, r, "/api/v1/createrequests", http.StatusSeeOther)
	}
}

//GetPendingRequests gets contract requests for user
func GetPendingRequests(w http.ResponseWriter, r *http.Request) {

}

//CreateRequest creates request new contract
func CreateRequest(w http.ResponseWriter, r *http.Request) {

}
