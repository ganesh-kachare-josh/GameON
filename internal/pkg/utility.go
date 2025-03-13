package pkg

import (
	"errors"
	"os"
	"strconv"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) 
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

func SendEmail(to string , subject string , body string) error{
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
		return err
	}

	smtp_server := os.Getenv("smtp_server")
	smtp_port := os.Getenv("smtp_port")
	smtp_user := os.Getenv("smtp_username")
	smtp_password := os.Getenv("smtp_password")

	port,err := strconv.Atoi(smtp_port)
		if err != nil {
			log.Println(err)
			return err
		}
	
	message := gomail.NewMessage()
	message.SetHeader("From", "GameON <" + smtp_user + ">") 
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body) 

	// Set up SMTP client
	dialer := gomail.NewDialer(smtp_server,port, smtp_user, smtp_password) 

	if err := dialer.DialAndSend(message); err != nil {
		log.Println(err)
		return errors.New("failed to send the email")
	}
	return nil 
}