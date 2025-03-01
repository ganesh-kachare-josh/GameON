package app

import (
	"net/http"

	"github.com/ganesh-kachare-josh/GameON/internal/app/auth"
	"github.com/ganesh-kachare-josh/GameON/internal/app/profile"
	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.CORS)
	// Request Routes
	router.HandleFunc("/request/{id}", middleware.AuthenticationMiddleware(request.GetRequestById(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/requests", middleware.AuthenticationMiddleware(request.GetAllRequests(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/request/{id}/participants",middleware.AuthenticationMiddleware(request.GetAllParticipants(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/request/{request_id}/accept",middleware.AuthenticationMiddleware(request.AcceptRequest(deps.RequestService))).Methods(http.MethodPost)

	router.HandleFunc("/request/{request_id}/confirm",middleware.AuthenticationMiddleware(request.ConfirmRequest(deps.RequestService))).Methods(http.MethodPost)

	router.HandleFunc("/request/{request_id}",middleware.AuthenticationMiddleware(request.DeleteRequest(deps.RequestService))).Methods(http.MethodDelete)

	router.HandleFunc("/participants/{pid}",middleware.AuthenticationMiddleware(request.RejectParticipant(deps.RequestService))).Methods(http.MethodDelete)

	router.HandleFunc("/request",middleware.AuthenticationMiddleware(request.CreateRequest(deps.RequestService))).Methods(http.MethodPost)

	// Profile Routes
	router.HandleFunc("/user/{user_id}",middleware.AuthenticationMiddleware(profile.GetUserById(deps.ProfileService))).Methods(http.MethodGet)

	// Authentication Routes.
	router.HandleFunc("/login", auth.Login(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/register", auth.Register(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/logout" ,middleware.AuthenticationMiddleware(auth.Logout(deps.AuthService))).Methods(http.MethodPost)
	router.HandleFunc("/islogin" ,auth.IsLogin(deps.AuthService)).Methods(http.MethodGet) 
	return router

}
