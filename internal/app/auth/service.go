package auth

import (
	"context"
	"github.com/ganesh-kachare-josh/GameON/internal/repository"
	"log"
)

type Service interface {
	Login(ctx context.Context, requestBody repository.Login) (LoginResponse, error)
	Register(ctx context.Context, requestBody RegisterData) (RegisterData, error)
}

type service struct {
	authRepo repository.RepoAuth
}

func (s *service) Login(ctx context.Context, requestBody repository.Login) (LoginResponse, error) {
	login, err := s.authRepo.Login(ctx, repository.Login(requestBody))
	if err != nil {
		log.Println(err)
		return LoginResponse{}, err
	}
	return LoginResponse(login) , err 
}

func (s *service) Register(ctx context.Context, requestBody RegisterData) (RegisterData, error) {
	register, err := s.authRepo.Register(ctx, repository.Register(requestBody))
	if err != nil {
		log.Println(err)
		return RegisterData{}, err
	}
	return RegisterData(register), nil
}

func NewService(authRepo repository.RepoAuth) Service {
	return &service{
		authRepo: authRepo,
	}
}
