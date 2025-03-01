package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(userID int) (string, error) {

	err := godotenv.Load("../.env")
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour).Unix() 

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp" : expirationTime , 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("jwt_secret_key") 
	jwtToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}


func VerifyToken(tokenString string) (*jwt.Token, error) {

	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwt_secret_key")), nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the verified token
	return token, nil
}
