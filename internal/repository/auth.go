package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type repoAuth struct {
	DB *sql.DB
}

type RepoAuth interface {
	Login(ctx context.Context, requestBody Login) (LoginResponse, error)
	Register(ctx context.Context, requestBody Register) (Register, error)
}

func NewAuthRepo(db *sql.DB) RepoAuth {
	return &repoAuth{
		DB: db,
	}
}

func (ra repoAuth) Login(ctx context.Context, requestBody Login) (LoginResponse, error) {
	db := sqlx.NewDb(ra.DB, "postgres")

	var login Login
	// originalPassword := requestBody.Password

	var response LoginResponse

	err := db.Get(&login, "SELECT id , email , name , password FROM users WHERE email = $1", requestBody.Email)
	if err != nil {
		log.Println(err)
		return LoginResponse{}, errors.New("incorrect email")
	}

	islogin := pkg.VerifyPassword(requestBody.Password, login.Password)
	if !islogin {
		return LoginResponse{}, errors.New("incorrect password")
	}

	// Generating JWT Token.
	tokenString, err := pkg.GenerateToken(login.Id)
	if err != nil {
		log.Println(err)
		return LoginResponse{}, err
	}

	response.Id = login.Id 
	response.Email=login.Email
	response.Name=login.Name
	response.Token = tokenString

	return response, nil
}

func (ra repoAuth) Register(ctx context.Context, requestBody Register) (Register, error) {
	db := sqlx.NewDb(ra.DB, "postgres")
	var register Register

	// Encryption of password
	originalPassword := requestBody.Password
	hashedPassword, err := pkg.HashPassword(originalPassword)
	if err != nil {
		log.Println(err)
		return Register{}, err
	}

	requestBody.Password = hashedPassword // Assigning hashed password to requestBody

	err = db.Get(&register, pkg.RegisterUserQuery, requestBody.Name, requestBody.Email, requestBody.Password, requestBody.Phone_Number, requestBody.Sport)

	if err != nil {
		log.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // 23505 is the error code for unique violation
				if pqErr.Message == "duplicate key value violates unique constraint \"users_email_unique\"" {
					return Register{}, errors.New("email already exists")
				}
				if pqErr.Message == "duplicate key value violates unique constraint \"users_phone_number_unique\"" {
					return Register{}, errors.New("phone number already exists")
				}
			}
		}

		return Register{}, err
	}
	register.Password = originalPassword // Reassigning original password to requestBody.
	return register, nil
}