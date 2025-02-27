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
	router.HandleFunc("/request/{id}/participants" , request.GetAllParticipants(deps.RequestService)).Methods(http.MethodGet)
	router.HandleFunc("/request/{request_id}/accept" , request.AcceptRequest(deps.RequestService)).Methods(http.MethodPost)
	router.HandleFunc("/request/{request_id}/confirm" , request.ConfirmRequest(deps.RequestService)).Methods(http.MethodPost)
	router.HandleFunc("/request/{request_id}" , request.DeleteRequest(deps.RequestService)).Methods(http.MethodDelete)
	router.HandleFunc("/participants/{pid}" , request.RejectParticipant(deps.RequestService)).Methods(http.MethodDelete)





	// Authentication. 
	router.HandleFunc("/login" , auth.Login(deps.AuthService)).Methods(http.MethodPost) 
	router.HandleFunc("/register" , auth.Register(deps.AuthService)).Methods(http.MethodPost) 
	return router

}