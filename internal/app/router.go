package app

import (
	"net/http"

	"github.com/ganesh-kachare-josh/GameON/internal/app/auth"
	"github.com/ganesh-kachare-josh/GameON/internal/app/profile"
	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {
	router := mux.NewRouter()

	// Request Routes
	router.HandleFunc("/request/{id}", pkg.AuthenticationMiddleware(request.GetRequestById(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/requests", pkg.AuthenticationMiddleware(request.GetAllRequests(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/request/{id}/participants", pkg.AuthenticationMiddleware(request.GetAllParticipants(deps.RequestService))).Methods(http.MethodGet)

	router.HandleFunc("/request/{request_id}/accept", pkg.AuthenticationMiddleware(request.AcceptRequest(deps.RequestService))).Methods(http.MethodPost)

	router.HandleFunc("/request/{request_id}/confirm", pkg.AuthenticationMiddleware(request.ConfirmRequest(deps.RequestService))).Methods(http.MethodPost)

	router.HandleFunc("/request/{request_id}", pkg.AuthenticationMiddleware(request.DeleteRequest(deps.RequestService))).Methods(http.MethodDelete)

	router.HandleFunc("/participants/{pid}", pkg.AuthenticationMiddleware(request.RejectParticipant(deps.RequestService))).Methods(http.MethodDelete)

	router.HandleFunc("/request", pkg.AuthenticationMiddleware(request.CreateRequest(deps.RequestService))).Methods(http.MethodPost)

	// Profile Routes
	router.HandleFunc("/user/{user_id}", pkg.AuthenticationMiddleware(profile.GetUserById(deps.ProfileService))).Methods(http.MethodGet)

	// Authentication Routes.
	router.HandleFunc("/login", auth.Login(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/register", auth.Register(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/logout" , pkg.AuthenticationMiddleware(auth.Logout(deps.AuthService))).Methods(http.MethodPost)
	return router

}
