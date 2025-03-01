package repository

import (
	"encoding/json"
	"time"
)

type Request struct {
	Id         int             `db:"id"`
	User_id    int             `db:"user_id"`
	Sport      json.RawMessage `db:"sport"`
	Location   string          `db:"location"`
	Time       time.Time       `db:"time"`
	CourtPrice float64         `db:"court_price"`
	Status     string          `db:"status"`
}

var Requests []Request

type Login struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type Register struct {
	Id           int             `db:"id"`
	Name         string          `db:"name"`
	Email        string          `db:"email"`
	Password     string          `db:"password"`
	Phone_Number string          `db:"phone_number"`
	Sport        json.RawMessage `db:"sports"`
	Created_at   string          `db:"created_at"`
}

type ParticipantData struct {
	Id     int    `db:"id"`
	UserId int    `db:"user_id"`
	Status string `db:"status"`
}

var Participants []ParticipantData

type AcceptRequestData struct {
	Id         int    `db:"id"`
	Request_id int    `db:"request_id"`
	User_id    int    `db:"user_id"`
	Status     string `db:"status"`
}

type AcceptRequestBody struct {
	Request_id int
	User_id    int
}

type UserData struct {
	Id           int             `db:"id"`
	Name         string          `db:"name"`
	Email        string          `db:"email"`
	Sports       json.RawMessage `db:"sports"`
	Phone_Number string          `db:"phone_number"`
}

type LoginResponse struct {
	LoginData Login
	Token     string
}

type LogoutResponse struct {
	Message string
}
