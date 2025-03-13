package pkg

import (
	"errors"
	"os"
	"time"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(userID int) (string, error) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return "", err
	}
	return jwtToken, nil
}


func VerifyToken(tokenString string) (*jwt.Token, error) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwt_secret_key")), nil
	})

	// Check for verification errors
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the verified token
	return token, nil
}

func GetUserIdFromToken(tokenString string) (int , error) {
	token , err := VerifyToken(tokenString) 
	if err != nil {
		log.Println(err)
		return 0 , err 
	} 

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user_id from the claims
		userID, ok := claims["user_id"].(float64) // user_id in claims should be a float64 type
		if ok {
			return int(userID), nil // Return user_id as an integer
		}
	}
	return 0 , err 
}
