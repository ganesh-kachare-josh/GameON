package profile

import (
	"context"

	"github.com/ganesh-kachare-josh/GameON/internal/repository"
)

type service struct {
	profileRepo repository.ProfileRepo
}

type Service interface {
	GetUserById(ctx context.Context, user_id int) (UserData, error)
	UpdateProfile(ctx context.Context , requestBody UserData)(UserData , error)
}

func (s *service) GetUserById(ctx context.Context, user_id int) (UserData, error) {
	response, err := s.profileRepo.GetUserById(ctx, user_id)
	if err != nil {
		return UserData{}, err
	}
	return UserData(response), nil
}

func (s *service) UpdateProfile(ctx context.Context , requestBody UserData)(UserData , error) {
	response , err :=  s.profileRepo.UpdateProfile(ctx , repository.UserData(requestBody))
	if err != nil {
		return UserData{} , err 
	}
	return UserData(response) , nil 
}

func NewService(profileRepo repository.ProfileRepo) Service {
	return &service{
		profileRepo: profileRepo,
	}
}
