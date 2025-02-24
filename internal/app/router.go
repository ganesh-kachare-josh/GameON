package app

import (
	"net/http"

	"github.com/ganesh-kachare-josh/GameON/internal/app/auth"
	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {
	router := mux.NewRouter()
	
	// Routes.
	router.HandleFunc("/request/{id}" , request.GetRequestById(deps.RequestService)).Methods(http.MethodGet)
	router.HandleFunc("/requests" , request.GetAllRequests(deps.RequestService)).Methods(http.MethodGet)

	// Authentication. 
	router.HandleFunc("/login" , auth.Login(deps.AuthService)).Methods(http.MethodPost) 
	router.HandleFunc("/register" , auth.Register(deps.AuthService)).Methods(http.MethodPost) 
	return router

} 