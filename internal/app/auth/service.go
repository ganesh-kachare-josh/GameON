package auth

import (
	"context"
	"github.com/ganesh-kachare-josh/GameON/internal/repository"
)

type Service interface {
	Login(ctx context.Context, requestBody repository.Login) (repository.LoginResponse, error)
	Register(ctx context.Context, requestBody RegisterData) (RegisterData, error)
	Logout(ctx context.Context) (repository.LogoutResponse)
}

type service struct {
	authRepo repository.RepoAuth
}

func (s *service) Login(ctx context.Context, requestBody repository.Login) (repository.LoginResponse, error) {
	login, err := s.authRepo.Login(ctx, repository.Login(requestBody))
	if err != nil {
		return repository.LoginResponse{}, err
	}
	return login, nil
}

func (s *service) Register(ctx context.Context, requestBody RegisterData) (RegisterData, error) {
	register, err := s.authRepo.Register(ctx, repository.Register(requestBody))
	if err != nil {
		return RegisterData{}, err
	}
	return RegisterData(register), nil
}

func (s *service) Logout(ctx context.Context) (repository.LogoutResponse) {
	return s.authRepo.Logout(ctx)
}

func NewService(authRepo repository.RepoAuth) Service {
	return &service{
		authRepo: authRepo,
	}
}
