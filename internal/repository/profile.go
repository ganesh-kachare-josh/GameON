package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
)
type profileRepo struct {
	DB *sql.DB
}

type ProfileRepo interface {
	GetUserById(ctx context.Context , user_id int) (UserData , error)
	UpdateProfile(ctx context.Context ,requestBody UserData) (UserData , error) 
}

func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &profileRepo{
		DB: db,
	}
}

func (rp profileRepo) GetUserById(ctx context.Context , user_id int) (UserData , error) {
	db := sqlx.NewDb(rp.DB , "postgres")
	
	var user UserData 

	err := db.Get(&user , pkg.GetUserByIdQuery , user_id) 
	if err != nil {
		log.Println(err)
		return UserData{} , fmt.Errorf("user Does Not Exist") 
	}
	return user , nil 
}

func (rp profileRepo) UpdateProfile(ctx context.Context ,requestBody UserData) (UserData , error) {
	db := sqlx.NewDb(rp.DB , "postgres") 
	
	var user UserData 

	err := db.Get(&user , pkg.UpdateUserByIdQuery, requestBody.Id , requestBody.Name , requestBody.Email , requestBody.Sports , requestBody.Phone_Number) 
	if err != nil {
		log.Println(err)
		return UserData{} , err
	}
	return user , nil 
}