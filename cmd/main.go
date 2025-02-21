package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	gameon "github.com/ganesh-kachare-josh/GameON"
)


func main() {
    ctx := context.Background()  
    
    err := godotenv.Load() 
    if err != nil {
        log.Fatal("Error in loading .env file")
        return 
    }

	sql , err := gameon.initDB(ctx) 

	if err != nil {
		fmt.Errorf("Error Occured while initializing database : %v", err)
		return 
	} 

	err = sql.Ping() 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database") 



	// Start the server 
	path := os.Getenv(db_host) + ":" + os.Getenv(db_port) 
	log.Fatal(http.ListenAndServe( path , nil ))
	fmt.Println("Server is running on port ")
}