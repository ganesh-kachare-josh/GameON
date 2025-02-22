package app

import (
	"net/http"

	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {
	router := mux.NewRouter()
	
	// Routes.
	router.HandleFunc("/request/{id}" , request.GetRequestById(deps.RequestService)).Methods(http.MethodGet)
	router.HandleFunc("/requests" , request.GetAllRequests(deps.RequestService)).Methods(http.MethodGet)
	return router

} 