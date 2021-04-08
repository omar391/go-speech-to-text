package utils

import (
	"log"
	"stt-service/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//generate a new outgoing JWT token
func GenerateJWT(user_id uint) (string, error) {
	// Set some claims
	claims := jwt.MapClaims{
		"uid": user_id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(conf.Config.JWT_SECRET)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

//verify the incoming token
func VerifyJWT(token string) bool {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Config.JWT_SECRET), nil
	})

	return err == nil && parsed.Valid
}
