package utils

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)


type loginStruct struct{
	Email string `json:"email"`
	UserId string `json:"userId"`
	Name string `json:"name"`
	AccessToken string `json:"accessToken"`
	jwt.RegisteredClaims
}
func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return  string(bytes),err
}


var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_KEY")
func GenerateToken(name,userId,email string)(string, error){
 claims := &loginStruct{
	Email: email,
	UserId: userId,
	Name: name,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: "oneofone",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	},
}
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 signedToken,err := token.SignedString([]byte(ACCESS_TOKEN_SECRET))
 if err != nil {
	return "", err
 }
 fmt.Println(signedToken)
 return signedToken, nil
}