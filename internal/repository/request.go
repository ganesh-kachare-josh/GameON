package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

const getRequestByIdQuery = "SELECT requests.id, requests.user_id, requests.sport, address.name, address.street,address.city, address.state, address.country,  requests.time,  requests.court_price, requests.status FROM requests JOIN address ON requests.address_id = address.id WHERE requests.id = $1;"

const GetAllRequestsQuery = "SELECT requests.id, requests.user_id, requests.sport, address.name, address.street,address.city, address.state, address.country,  requests.time,  requests.court_price, requests.status FROM requests JOIN address ON requests.address_id = address.id;"


type repoPerson struct {
	DB *sql.DB
}

type RepoPerson interface {
     GetRequestById(ctx context.Context , request_id int ) (Request , error) 
	 GetAllRequests(ctx context.Context) ([]Request , error) 
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