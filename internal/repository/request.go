package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
	"github.com/jmoiron/sqlx"
)

type repoPerson struct {
	DB *sql.DB
}

type RepoPerson interface {
     GetRequestById(ctx context.Context , request_id int ) (Request , error) 
	 GetAllRequests(ctx context.Context) ([]Request , error) 
	 GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData)
	 AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error)
	 ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error)
	 DeleteRequest(ctx context.Context , request_id int) (sql.Result , error)
	 RejectParticipant(ctx context.Context , participant_id int) (error)
	 CreateRequest(ctx context.Context , requestBody Request) (Request , error)
}

func NewRepo(db *sql.DB) (RepoPerson) {
	return &repoPerson {
		DB: db,
	}
}


func (rp repoPerson) GetRequestById(ctx context.Context , request_id int) (Request , error) {
  	db := sqlx.NewDb(rp.DB , "postgres") 
	var request Request 
	
	err := db.Get(&request , pkg.GetRequestByIdQuery , request_id)  
	if err != nil {
		return Request{} , err 
	}

	return request , nil 

}

func (rp repoPerson) GetAllRequests(ctx context.Context) ([]Request , error) {
	db := sqlx.NewDb(rp.DB , "postgres") 
	var requests []Request 

	err := db.Select(&requests ,pkg.GetAllRequestsQuery)  
	if err != nil {
		return []Request{} , err 
	}
    
	return requests , nil 
}

func (rp repoPerson) GetAllParticipants(ctx context.Context , request_id int) ([]ParticipantData) {
	db := sqlx.NewDb(rp.DB, "postgres")

	var participants []ParticipantData 

	err := db.Select(&participants ,pkg.GetAllParticipantsQuery , request_id)
	if err != nil {
		return []ParticipantData{}
	}
	return participants
}

func (rp repoPerson) AcceptRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error) {
	db := sqlx.NewDb(rp.DB, "postgres") 

	var data AcceptRequestData 

	err := db.Get(&data , pkg.AcceptRequestQuery , requestBody.Request_id, requestBody.User_id , "Pending")
	if err != nil {
		return AcceptRequestData{} , err
	}

	var userId int 
	err = db.Get(&userId , pkg.GetUserIdByRequestId , requestBody.Request_id)
	if err != nil {
		return AcceptRequestData{} , nil 
	} 

	var email string 
	err = db.Get(&email , pkg.GetEmailById , userId) 
	if err != nil {
		return AcceptRequestData{} , err 
	}

	var creatorName string 
	var ParticipantName string  
	var sport json.RawMessage 

	err = db.Get(&creatorName , pkg.GetNameByIdQuery , userId) 
	if err != nil {
		return AcceptRequestData{} , errors.New("user doesn't exist") 
	}

	err = db.Get(&ParticipantName , pkg.GetNameByIdQuery , requestBody.User_id) 
	if err != nil {
		return AcceptRequestData{} , errors.New("participant doesn't exist") 
	}

	err = db.Get(&sport , pkg.GetSportByRequestId , requestBody.Request_id) 
	if err != nil {
		return AcceptRequestData{} , errors.New("participant doesn't exist") 
	}
	var sportMap map[string]string 
	err = json.Unmarshal(sport , &sportMap) 
	if err != nil {
		return AcceptRequestData{} , errors.New("error unmarshaling sport")
	}

	var game string 
	for key := range sportMap {
		game = key 
		break
	}


	err = pkg.SendEmail(email, 
		fmt.Sprintf("🎉 Game On! %v Your Request Was Accepted!" , creatorName), 
		fmt.Sprintf("Great news! %v has accepted your game request to play %v. Get ready to jump into action and enjoy the thrill! 🚀\n\nLog in now to check the details and start gaming!\n\nHappy Gaming! 🎮" , ParticipantName , game),
	)
	
	if err != nil {
		return AcceptRequestData{} , err 
	}

	return data , nil 
}

func (rp repoPerson) ConfirmRequest(ctx context.Context , requestBody AcceptRequestBody) (AcceptRequestData , error) {
	db := sqlx.NewDb(rp.DB, "postgres") 

	var data AcceptRequestData 

	err := db.Get(&data , pkg.ConfirmRequestQuery , requestBody.Request_id , requestBody.User_id)
	if err != nil {
		return AcceptRequestData{} , err 
	}

	var email string 
	err = db.Get(&email , pkg.GetEmailById , requestBody.User_id) 
	if err != nil {
		return AcceptRequestData{} , err 
	} 

	var userId int 
	err = db.Get(&userId , pkg.GetUserIdByRequestId , requestBody.Request_id)
	if err != nil {
		return AcceptRequestData{} , nil 
	}

	var creatorName string 
	var sport json.RawMessage 

	err = db.Get(&creatorName , pkg.GetNameByIdQuery , userId) 
	if err != nil {
		return AcceptRequestData{} , errors.New("user doesn't exist") 
	}

	err = db.Get(&sport , pkg.GetSportByRequestId , requestBody.Request_id) 
	if err != nil {
		return AcceptRequestData{} , errors.New("participant doesn't exist") 
	}
	var sportMap map[string]string 
	err = json.Unmarshal(sport , &sportMap) 
	if err != nil {
		return AcceptRequestData{} , errors.New("error unmarshaling sport")
	}

	var game string 
	for key := range sportMap {
		game = key 
		break
	}

	
	err = pkg.SendEmail(email, 
		"✅ You're In! Join Request Confirmed", 
		fmt.Sprintf("Congratulations! %v has accepted your join request to play %v. You're now part of the squad! 🔥\n\nPrepare yourself, gear up, and get ready for an epic gaming session.\n\nSee you in the game! 🎮",creatorName , game),
	)
	if err != nil {
		return AcceptRequestData{} , err 
	}

	return data , nil	
}

func (rp  repoPerson) DeleteRequest(ctx context.Context , request_id int) (sql.Result , error) {
	db := sqlx.NewDb(rp.DB, "postgres")
	
	result , err := db.Exec(pkg.DeleteRequestQuery , request_id) 
	if err != nil {
		return result , fmt.Errorf("failed to delete item: %v", err)	
	}
    
	return result , nil 
}

func (rp repoPerson) RejectParticipant(ctx context.Context , participant_id int) (error) {
	db := sqlx.NewDb(rp.DB, "postgres")
	
	var user_id int 
	err := db.Get(&user_id , "SELECT user_id from participants WHERE id = $1" , participant_id) 
	if err != nil {
		return err 
	}	

	var email string 
	err = db.Get(&email , pkg.GetEmailById , user_id) 
	if err != nil {
		return err
	}

	var request_id int 
	err = db.Get(&request_id , "SELECT request_id from participants WHERE id = $1" , participant_id) 
	if err != nil {
		return err 
	}

	var creator_id int 
	err = db.Get(&creator_id , "SELECT user_id from requests WHERE id = $1" , request_id) 
	if err != nil {
		return err 
	}

	var creatorName string 
	var sport json.RawMessage  

	err = db.Get(&creatorName , pkg.GetNameByIdQuery , creator_id) 
	if err != nil {
		return err 
	}

	err = db.Get(&sport , pkg.GetSportByRequestId , request_id) 
	if err != nil {
		return err  
	}
	var sportMap map[string]string 
	err = json.Unmarshal(sport , &sportMap) 
	if err != nil {
		return err 
	}

	var game string 
	for key := range sportMap {
		game = key 
		break
	}


	_, err = db.Exec(pkg.RejectParticipantQuery , participant_id)  
	if err != nil {
		return fmt.Errorf("failed to delete item: %v", err)	
	}

	err = pkg.SendEmail(email, 
		fmt.Sprintf("❌ Oops! Join Request Rejected by %v" , creatorName), 
		fmt.Sprintf("Hey there, unfortunately, your request to join the game %v was not accepted this time. But don’t worry, new opportunities are always around the corner! 🌟\n\nKeep exploring, find another game, and show them what they’re missing!\n\nBetter luck next time! 🎮",game), 
	)
	if err != nil {
		return err 
	}
	

	return nil 
}

func (rp repoPerson) CreateRequest(ctx context.Context , requestBody Request) (Request , error) {
	db := sqlx.NewDb(rp.DB , "postgres") 

	var responseBody Request 
	err := db.Get(&responseBody , pkg.CreateRequestQuery , requestBody.User_id , requestBody.Sport , requestBody.Location , requestBody.Time , requestBody.CourtPrice)
	if err != nil {
		return Request{} , err 
	}
	return responseBody , nil 
}