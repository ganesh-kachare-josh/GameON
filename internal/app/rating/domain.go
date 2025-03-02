package rating

import (
	"encoding/json"
	"time"
)

type RatingResponse struct {
	Message string `json:"message"`
}

type RatingRequestBody struct {
	GivenBy    int    `json:"given_by"`
	GivenTo    int    `json:"given_to"`
	Request_id int    `json:"request_id"`
	Rating     int    `json:"rating"`
	Feedback   string `json:"feedback"`
}

type RatingUserIdResponse struct {
	Id         int             `json:"id"`
	Given_by   int             `json:"given_by"`
	Name       string          `json:"name"`
	Given_to   int             `json:"given_to"`
	Request_id int             `json:"request_id"`
	Sport      json.RawMessage `json:"sport"`
	Rating     int             `json:"rating"`
	Feedback   string          `json:"feedback"`
	Created_at time.Time       `json:"created_at"`
}
