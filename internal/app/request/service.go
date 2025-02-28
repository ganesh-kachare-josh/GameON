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
	GetAllRequests(ctx context.Context) ([]repository.Request , error)
	GetAllParticipants(ctx context.Context , request_id int)([]repository.ParticipantData)
	AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
	ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
	DeleteRequest(ctx context.Context , request_id int) (sql.Result , error)
	RejectParticipant(ctx context.Context , participant_id int) (error)
	CreateRequest(ctx context.Context , requestBody Request) (repository.Request , error) 
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

func (s *service) GetAllParticipants(ctx context.Context , request_id int) ([]repository.ParticipantData) {
		participants := s.requestRepo.GetAllParticipants(ctx , request_id)
		return participants
}

func (s *service) AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData) {
	response := s.requestRepo.AcceptRequest(ctx , repository.AcceptRequestBody(requestBody)) 
	return AcceptRequestData(response)
}

func (s *service) ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData) {
	response := s.requestRepo.ConfirmRequest(ctx , repository.AcceptRequestBody(requestBody))
	return AcceptRequestData(response)
}

func (s *service) DeleteRequest(ctx context.Context , request_id int) (sql.Result , error) {
	return s.requestRepo.DeleteRequest(ctx , request_id)
}

func (s *service) RejectParticipant(ctx context.Context , participant_id int) (error) {
	return s.requestRepo.RejectParticipant(ctx , participant_id)
}

func (s *service) CreateRequest(ctx context.Context , requestBody Request) (repository.Request , error) {
	return s.requestRepo.CreateRequest(ctx , repository.Request(requestBody))
}

func NewService(requestRepo repository.RepoPerson) (Service) {
      return &service{
		requestRepo: requestRepo,	
	  }
}