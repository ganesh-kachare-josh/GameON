package repository

import (
	"encoding/json"
	"time"
)

type Request struct {
	Id         int             `db:"id"`
	User_id    int             `db:"user_id"`
	Sport      json.RawMessage `db:"sport"`
	Name       string          `db:"name"`
	Street     string          `db:"street"`
	City       string          `db:"city"`
	State      string          `db:"state"`
	Country    string          `db:"country"`
	Time       time.Time       `db:"time"`
	CourtPrice float64         `db:"court_price"`
	Status     string          `db:"status"`
}

var Requests []Request

type Login struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

type Register struct {
	Name         string          `db:"name"`
	Email        string          `db:"email"`
	Password     string          `db:"password"`
	Phone_Number string          `db:"phone_number"`
	Sport        json.RawMessage `db:"sports"`
	Created_at   string          `db:"created_at"`
}
