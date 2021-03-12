package router

import (
	"go_Apprenticeship/handlers"
	"net/http"
)

//Router is the start point of the application
func Router() {
	http.HandleFunc("/api/v1/users", handlers.Index)
	http.HandleFunc("/api/v1/users/login", handlers.LoginUser)
	http.HandleFunc("/api/v1/users/signup", handlers.SignupUser)
	http.HandleFunc("/api/v1/users/logout", handlers.LogoutUser)
	http.HandleFunc("/api/v1/users/main", handlers.Home)
	http.HandleFunc("/api/v1/requests", handlers.Requests)
	http.HandleFunc("/api/v1/createrequests", handlers.CreateRequest)
	http.HandleFunc("/api/v1/getrequests", handlers.GetPendingRequests)

	http.ListenAndServe(":8080", nil)
}
