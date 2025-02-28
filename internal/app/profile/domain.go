package profile

import (
	"encoding/json"
)

type UserData struct {
	Id           int             `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Sports       json.RawMessage `json:"sports"`
	Phone_Number string          `json:"phone_number"`
}
