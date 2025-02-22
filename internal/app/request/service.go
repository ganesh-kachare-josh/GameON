package request

import (
	"context" 
	"github.com/ganesh-kachare-josh/GameON/internal/repository" 
) 

type service struct {
	requestRepo repository.RepoPerson 
}

type Service interface {
	GetRequestById(ctx context.Context , request_id int) (Request , error) 
}

func (s *service ) GetRequestById(ctx context.Context , request_id int) (Request , error) {
	    request , err := s.requestRepo.GetRequestById(ctx , request_id) 
		if err != nil {
			return Request{} , err 
		}
		return Request(request) , nil 
}

func NewService(requestRepo repository.RepoPerson) (Service) {
      return &service{
		requestRepo: requestRepo,
	  }
} 