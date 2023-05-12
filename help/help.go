package help

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

type UserClaims struct {
	Identity string
	Name     string
	jwt.StandardClaims
}

var key = []byte("gin-gorm-key")

func GetMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(identity, name string) (string, error) {
	userClaims := &UserClaims{
		Identity:       "identity",
		Name:           "name",
		StandardClaims: jwt.StandardClaims{},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {

	userClaims := new(UserClaims)

	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims.Valid {
		return userClaims, nil
	} else {
		return nil, fmt.Errorf("an error occurred on parsing token")
	}
}

func GetUUID() string {
	return uuid.New().String()
}

func GetRandom() string {
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
