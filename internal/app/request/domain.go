package request

import (
	"encoding/json"
	"time"
)

type Request struct {
	Id         int             `json:"Id"`
	User_id    int             `json:"User_id"`
	Sport      json.RawMessage `json:"Sport"`
	Name       string          `json:"Name"`
	Street     string          `json:"Street"`
	City       string          `json:"City"`
	State      string          `json:"State"`
	Country    string          `json:"Country"`
	Time       time.Time       `json:"Time"`
	CourtPrice float64         `json:"CourtPrice"`
	Status     string          `json:"Status"`
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
