package repository

import (
	"encoding/json"
	"time"
) 


type Request struct {
	Id          int                 `db:"id"`
	User_id     int                 `db:"user_id"`
	Sport       json.RawMessage   	`db:"sport"` 
	Name 		string 				`db:"name"`
	Street 		string 				`db:"street"`
	City 		string				`db:"city"`
	State		string 				`db:"state"`
	Country 	string 				`db:"country"`
	Time        time.Time           `db:"time"` 
	CourtPrice  float64             `db:"court_price"`
	Status      string              `db:"status"`	
}




