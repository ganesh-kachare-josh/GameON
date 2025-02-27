package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const GetUserByIdQuery = "SELECT id , name , email , sports , phone_number FROM users WHERE id = $1"

type profileRepo struct {
	DB *sql.DB
}

type ProfileRepo interface {
	GetUserById(ctx context.Context , user_id int) (UserData , error)
}

func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &profileRepo{
		DB: db,
	}
}

func (rp profileRepo) GetUserById(ctx context.Context , user_id int) (UserData , error) {
	db := sqlx.NewDb(rp.DB , "postgres")
	
	var user UserData 

	err := db.Get(&user , GetUserByIdQuery , user_id) 
	if err != nil {
		return UserData{} , fmt.Errorf("user Does Not Exist") 
	}
	return user , nil 
} 