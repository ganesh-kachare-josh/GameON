package app

import (
	"database/sql"

	"github.com/ganesh-kachare-josh/GameON/internal/app/request"
	"github.com/ganesh-kachare-josh/GameON/internal/repository"
)

type Dependencies struct {
	RequestService request.Service
}

func NewServices(db *sql.DB) (Dependencies) {
	requestRepo := repository.NewRepo(db) 
	requestService := request.NewService(requestRepo) 

	return Dependencies{
		RequestService: requestService,
	}
}

