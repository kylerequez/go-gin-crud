package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaims struct {
	Authorized string
	Email      string
	Expiration time.Time
}

func GetSecretKey() []byte {
	return secretKey
}

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["expiration"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(GetSecretKey())

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(t string) (token *jwt.Token, err error) {
	resultToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return secretKey, nil
	})

	fmt.Println("Token value is: ", resultToken)

	if err != nil {
		fmt.Println("There was an error in the validation: ", err)
		return nil, err
	}
	return resultToken, nil
}
