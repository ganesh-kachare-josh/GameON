package app

import (
	"database/sql"

	"github.com/ganesh-kachare-josh/GameON/internal/app/auth"
	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/ganesh-kachare-josh/GameON/internal/repository"
)

type Dependencies struct {
	RequestService request.Service
	AuthService auth.Service
}

func NewServices(db *sql.DB) (Dependencies) {
	requestRepo := repository.NewRepo(db) 
	requestService := request.NewService(requestRepo) 

	authRepo := repository.NewAuthRepo(db)
	authService := auth.NewService(authRepo) 


	return Dependencies{
		RequestService: requestService,
		AuthService: authService,
	}
}

