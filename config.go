package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func InitDB(ctx context.Context) (*sql.DB, error) {

	// This loads the .env file
	err := godotenv.Load("../.env") 
	if err != nil {
		return nil, errors.New("error in loading .env file")
	}

	dbConfig := databaseConfig{
		host:     os.Getenv("db_host"),
		port:     os.Getenv("db_port"),
		user:     os.Getenv("db_user"),
		password: os.Getenv("db_password"),
		dbname:   os.Getenv("dbname"),
	}

	databaseString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.dbname,
	)
	// Opening the db connection.
	db, err := sql.Open("postgres", databaseString)

	if err != nil {
		return nil, fmt.Errorf("error in opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Database connection established successfully")
	return db, nil

}
