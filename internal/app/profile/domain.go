package profile

import (
	"encoding/json"
	"time"
)

type UserData struct {
	Id           int             `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Sports       json.RawMessage `json:"sports"`
	Phone_Number string          `json:"phone_number"`
	Created_at   time.Time       `json:"created_at"`
}
