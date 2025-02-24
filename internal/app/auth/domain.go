package auth

import (
	"encoding/json"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterData struct {
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	Phone_Number string          `json:"phone_number"`
	Sport        json.RawMessage `json:"sports"`
	Created_at   string          `json:"created_at"`
}
