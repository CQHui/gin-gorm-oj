package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strconv"
)

type UserClaims struct {
	Identity string
	Name     string
	IsAdmin  int
	jwt.StandardClaims
}

var key = []byte("gin-gorm-key")

func GetMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(identity, name string, isAdmin int) (string, error) {
	userClaims := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        isAdmin,
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

func CodeSave(code []byte) (string, error) {
	dirName := GetUUID()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = f.Write(code)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	return path, nil
}
