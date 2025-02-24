package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const registerUserQuery = "INSERT INTO users (name , email , password , phone_number , sports , created_at) VALUES ($1 , $2 , $3 , $4 , $5 , NOW()) RETURNING name , email , password , phone_number , sports , created_at"

type repoAuth struct {
	DB *sql.DB
}

type RepoAuth interface {
	Login(ctx context.Context, requestBody Login) (Login, error)
	Register(ctx context.Context, requestBody Register) (Register, error)
}

func NewAuthRepo(db *sql.DB) RepoAuth {
	return &repoAuth{
		DB: db,
	}
}

func (ra repoAuth) Login(ctx context.Context, requestBody Login) (Login, error) {
	db := sqlx.NewDb(ra.DB, "postgres")
	var login Login
	originalPassword := requestBody.Password

	err := db.Get(&login, "SELECT email , password FROM users WHERE email = $1", requestBody.Email)
	if err != nil {
		return Login{}, errors.New("incorrect email")
	}

	islogin := pkg.VerifyPassword(requestBody.Password, login.Password)

	if !islogin {
		return Login{}, errors.New("incorrect password")
	}

	login.Password = originalPassword
	return login, nil
}

func (ra repoAuth) Register(ctx context.Context, requestBody Register) (Register, error) {
	db := sqlx.NewDb(ra.DB, "postgres")
	var register Register

	// Encryption of password
	originalPassword := requestBody.Password
	hashedPassword, err := pkg.HashPassword(originalPassword)
	if err != nil {
		return Register{}, err
	}

	requestBody.Password = hashedPassword // Assigning hashed password to requestBody

	err = db.Get(&register, registerUserQuery, requestBody.Name, requestBody.Email, requestBody.Password, requestBody.Phone_Number, requestBody.Sport)

	if err != nil {

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

	requestBody.Password = originalPassword // Reassigning original password to requestBody.
	return register, nil
}
