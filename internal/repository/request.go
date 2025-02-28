package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

const getRequestByIdQuery = "SELECT id , user_id , sport , location , time , court_price , status FROM requests WHERE id = $1"

const GetAllRequestsQuery = "SELECT id , user_id , sport , location , time , court_price , status FROM requests"

const GetAllParticipantsQuery = "SELECT id , user_id , status FROM participants WHERE request_id = $1"

const AcceptRequestQuery = "INSERT INTO participants (request_id , user_id , status) VALUES($1 ,$2 ,$3) RETURNING *"

const ConfirmRequestQuery = "UPDATE participants SET status = REPLACE(status , 'Pending' , 'Confirmed') WHERE request_id = $1 AND user_id = $2 RETURNING *"

const DeleteRequestQuery = "DELETE FROM requests WHERE id = $1"

const RejectParticipantQuery = "DELETE FROM participants WHERE id = $1"

type repoPerson struct {
	DB *sql.DB
}

type RepoPerson interface {
     GetRequestById(ctx context.Context , request_id int ) (Request , error) 
	 GetAllRequests(ctx context.Context) ([]Request , error) 
	 GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData)
	 AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
	 ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData)
	 DeleteRequest(ctx context.Context , request_id int) (sql.Result , error)
	 RejectParticipant(ctx context.Context , participant_id int) (error)
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

func (rp  repoPerson) DeleteRequest(ctx context.Context , request_id int) (sql.Result , error) {
	db := sqlx.NewDb(rp.DB, "postgres")
	
	result , err := db.Exec(DeleteRequestQuery , request_id) 
	if err != nil {
		return result , fmt.Errorf("failed to delete item: %v", err)	
	}
    
	return result , nil 
}

func (rp repoPerson) RejectParticipant(ctx context.Context , participant_id int) (error) {
	db := sqlx.NewDb(rp.DB, "postgres")
	
	_, err := db.Exec(RejectParticipantQuery , participant_id)  
	if err != nil {
		return fmt.Errorf("failed to delete item: %v", err)	
	}
	return nil 
}