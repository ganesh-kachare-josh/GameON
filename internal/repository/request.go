package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

const getRequestByIdQuery = "SELECT requests.id, requests.user_id, requests.sport, address.name, address.street,address.city, address.state, address.country,  requests.time,  requests.court_price, requests.status FROM requests JOIN address ON requests.address_id = address.id WHERE requests.id = $1;"

const GetAllRequestsQuery = "SELECT requests.id, requests.user_id, requests.sport, address.name, address.street,address.city, address.state, address.country,  requests.time,  requests.court_price, requests.status FROM requests JOIN address ON requests.address_id = address.id;"

const GetAllParticipantsQuery = "SELECT id , user_id , status FROM participants WHERE request_id = $1"

const AcceptRequestQuery = "INSERT INTO participants (request_id , user_id , status) VALUES($1 ,$2 ,$3) RETURNING *"

const ConfirmRequestQuery = "UPDATE participants SET status = REPLACE(status , 'Pending' , 'Confirmed') WHERE request_id = $1 AND user_id = $2 RETURNING *"

type repoPerson struct {
	DB *sql.DB
}

type RepoPerson interface {
     GetRequestById(ctx context.Context , request_id int ) (Request , error) 
	 GetAllRequests(ctx context.Context) ([]Request , error) 
	 GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData)
	 AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
	 ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
}

func NewRepo(db *sql.DB) (RepoPerson) {
	return &repoPerson {
		DB: db,
	}
}


func (rp repoPerson) GetRequestById(ctx context.Context , request_id int) (Request , error) {
  	db := sqlx.NewDb(rp.DB , "postgres") 
	var request Request 
	
	err := db.Get(&request , getRequestByIdQuery , request_id)  
	if err != nil {
		return Request{} , err 
	}

	return request , nil 

}

func (rp repoPerson) GetAllRequests(ctx context.Context) ([]Request , error) {
	db := sqlx.NewDb(rp.DB , "postgres") 
	var requests []Request 

	err := db.Select(&requests , GetAllRequestsQuery)  
	if err != nil {
		return []Request{} , err 
	}
    
	return requests , nil 
}

func (rp repoPerson) GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData) {
	db := sqlx.NewDb(rp.DB, "postgres")

	var participants []ParticipantData 

	err := db.Select(&participants , GetAllParticipantsQuery , request_id)
	if err != nil {
		return []ParticipantData{}
	}
	return participants
}

func (rp repoPerson) AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData) {
	db := sqlx.NewDb(rp.DB, "postgres") 

	var data AcceptRequestData 

	err := db.Get(&data , AcceptRequestQuery , requestBody.Request_id, requestBody.User_id , "Pending")
	if err != nil {
		return AcceptRequestData{}
	}
	return data
}

func (rp repoPerson) ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData) {
	db := sqlx.NewDb(rp.DB, "postgres") 

	var data AcceptRequestData 

	err := db.Get(&data , ConfirmRequestQuery , requestBody.Request_id , requestBody.User_id)
	if err != nil {
		return AcceptRequestData{}
	}
	return data
}