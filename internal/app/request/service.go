package request

import (
	"context"
	"database/sql"

	"github.com/ganesh-kachare-josh/GameON/internal/repository"
) 

type service struct {
	requestRepo repository.RepoPerson 
}

type Service interface {
	GetRequestById(ctx context.Context , request_id int) (Request , error) 
	GetAllRequests(ctx context.Context) ([]Request , error)
	GetAllParticipants(ctx context.Context , request_id int)([]ParticipantData)
	AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error)
	ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error) 
	DeleteRequest(ctx context.Context , request_id int) (sql.Result , error)
	RejectParticipant(ctx context.Context , participant_id int) (error)
	CreateRequest(ctx context.Context , requestBody Request) (Request , error) 
}

func (s *service ) GetRequestById(ctx context.Context , request_id int) (Request , error) {
	    request , err := s.requestRepo.GetRequestById(ctx , request_id) 
		if err != nil {
			return Request{} , err 
		}
		return Request(request) , nil 
}

func (s *service ) GetAllRequests(ctx context.Context) ([]Request, error) {
	    requests , err := s.requestRepo.GetAllRequests(ctx) 
		if err != nil {
			return []Request{} , err 
		}
		// var Requests []Request
		var result []Request
    	for _, r := range requests {
        	result = append(result, Request(r))
    	}
		return  result, nil 
}

func (s *service) GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData) {
		participants := s.requestRepo.GetAllParticipants(ctx , request_id)
		var result []ParticipantData
    	for _, r := range participants {
        	result = append(result, ParticipantData(r))
    	}
		return result
}

func (s *service) AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error) {
	response , err := s.requestRepo.AcceptRequest(ctx , repository.AcceptRequestBody(requestBody))
	if err != nil {
		return AcceptRequestData{} , err 
	} 
	return AcceptRequestData(response) , nil 
}

func (s *service) ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error) {
	response , err  := s.requestRepo.ConfirmRequest(ctx , repository.AcceptRequestBody(requestBody))
	if err != nil {
		return AcceptRequestData{} , err 
	}
	return AcceptRequestData(response) , nil 
}

func (s *service) DeleteRequest(ctx context.Context , request_id int) (sql.Result , error) {
	return s.requestRepo.DeleteRequest(ctx , request_id)
}

func (s *service) RejectParticipant(ctx context.Context , participant_id int) (error) {
	return s.requestRepo.RejectParticipant(ctx , participant_id)
}

func (s *service) CreateRequest(ctx context.Context , requestBody Request) (Request, error) {
	response , err := s.requestRepo.CreateRequest(ctx , repository.Request(requestBody))
	if err != nil {
		return Request{} , err 
	}
	return Request(response) , nil 
}

func NewService(requestRepo repository.RepoPerson) (Service) {
      return &service{
		requestRepo: requestRepo,	
	  }
}