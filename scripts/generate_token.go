package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("my-secret") // match JWT_SECRET from .env

	claims := jwt.MapClaims{
		"sub": "user-123",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bearer " + signed)
}
