package repository 

import (
	"context"
	"database/sql"
	"log"

	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
	"github.com/jmoiron/sqlx"
)

type repoRating struct {
	DB *sql.DB
}

type RepoRating interface {
	GiveRating(ctx context.Context , requestBody RatingRequestBody) (RatingResponse , error)
	GetRatingByUserId(ctx context.Context , user_id int) ([]RatingUserIdResponse , error)
}

func (rr repoRating) GiveRating(ctx context.Context , requestBody RatingRequestBody) (RatingResponse , error) {
	db := sqlx.NewDb(rr.DB, "postgres")  

	_ , err := db.Exec(pkg.GiveRatingQuery ,requestBody.GivenBy , requestBody.GivenTo , requestBody.Request_id , requestBody.Rating,requestBody.Feedback) 
	if err != nil {
		log.Println(err)
		return RatingResponse{} , err  
	}

	var response RatingResponse 
	response.Message = "Rating Given Successfully." 

	return response , nil 
}

func (rr repoRating) GetRatingByUserId(ctx context.Context , user_id int) ([]RatingUserIdResponse , error) {
	db := sqlx.NewDb(rr.DB, "postgres")  

	var response []RatingUserIdResponse 
	err := db.Select(&response , pkg.GetRatingByUserIdQuery , user_id) 
	if err != nil {
		log.Println(err)
		return []RatingUserIdResponse{} , err 
	}
	return response , nil 
}

func NewRatingRepo(db *sql.DB) RepoRating {
	return &repoRating{
		DB: db,
	}
}