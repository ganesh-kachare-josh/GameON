package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	database "github.com/ganesh-kachare-josh/GameON"
	"github.com/ganesh-kachare-josh/GameON/internal/app"
)


func main() {

    ctx := context.Background()  
    

    err := godotenv.Load("../.env") 
    if err != nil {
        log.Fatalf("Error in loading .env file %v" , err)
        return 
    }
    
	// Setting the database.
	sql , err := database.InitDB(ctx) 
	if err != nil {
		log.Fatalf("error Occured while initializing database : %v", err)
		return 
	} 
	
	// Injecting dependencies.
    services := app.NewServices(sql) 
	router := app.NewRouter(services) 

	srv := &http.Server {
		Addr: ":" + os.Getenv("port"),
		Handler: router,
	}

	fmt.Printf("Server is running on port : %v" , os.Getenv("port")) 
	log.Fatal(srv.ListenAndServe())
}