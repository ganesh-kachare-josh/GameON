package request

import (
	"encoding/json"
	"time"
)

type Request struct {
	Id           int             `json:"id"`
	User_id      int             `json:"user_id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Phone_Number string          `json:"phone_number"`
	Sport        json.RawMessage `json:"sport"`
	Location     string          `json:"location"`
	Time         time.Time       `json:"time"`
	CourtPrice   float64         `json:"court_price"`
	Status       string          `json:"status"`
}

type AcceptRequestData struct {
	Id         int    `json:"id"`
	Request_id int    `json:"request_id"`
	User_id    int    `json:"user_id"`
	Status     string `json:"status"`
}

type AcceptRequestBody struct {
	Request_id int `json:"request_id"`
	User_id    int `json:"user_id"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type ParticipantData struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
