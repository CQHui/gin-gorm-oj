package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type userClaims struct {
	Identity string
	Name     string
	jwt.StandardClaims
}

var key = []byte("gin-gorm-key")

func TestGenerateToken(t *testing.T) {
	userClaims := &userClaims{
		Identity:       "identity",
		Name:           "name",
		StandardClaims: jwt.StandardClaims{},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := claims.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(tokenString)
}

func TestParseToken(t *testing.T) {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZGVudGl0eSI6ImlkZW50aXR5IiwiTmFtZSI6Im5hbWUifQ.IkJzmoM44fsnczL9GeNtYxkrJJgzvogzyI1b2Bh3SHg"

	userClaims := new(userClaims)

	claims, err := jwt.ParseWithClaims(token, userClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaims)
	}
}
