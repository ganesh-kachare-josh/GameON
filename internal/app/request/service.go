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
	GetAllRequests(ctx context.Context) ([]repository.Request , error)
}

func (s *service ) GetRequestById(ctx context.Context , request_id int) (Request , error) {
	    request , err := s.requestRepo.GetRequestById(ctx , request_id) 
		if err != nil {
			return Request{} , err 
		}
		return Request(request) , nil 
}

func (s *service ) GetAllRequests(ctx context.Context) ([]repository.Request , error) {
	    requests , err := s.requestRepo.GetAllRequests(ctx) 
		if err != nil {
			return []repository.Request{} , err 
		}
		return  requests, nil 
}


func NewService(requestRepo repository.RepoPerson) (Service) {
      return &service{
		requestRepo: requestRepo,
	  }
} 