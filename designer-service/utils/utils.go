package utils

import (
	"errors"
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
var REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_KEY")

func GenerateToken(name,userId,email string)(string, error){
	
 claims := &loginStruct{
	Email: email,
	UserId: userId,
	Name: name,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: "oneofone",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
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


func GenerateAcessAndRefreshToken (id,email,name string)(string,string,error){
	 accessClaims := &loginStruct{
	Email: email,
	UserId: id,
	Name: name,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: "oneofone",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	},
}
  	 refreshClaims := &loginStruct{
	Email: email,
	UserId: id,
	Name: name,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: "oneofone",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
	},
}

accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
access,err := accessToken.SignedString([]byte(REFRESH_TOKEN_SECRET))


if err!=nil {
	return "","",err
}
 refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
 refresh,err := refreshToken.SignedString([]byte(ACCESS_TOKEN_SECRET))
 if err != nil {
	return "","",err
 }
 return access,refresh, nil
}


func VerifyToken(token string) (*loginStruct,error){
  claims := &loginStruct{}

	result,err := jwt.ParseWithClaims(token, claims, func (token *jwt.Token) (interface{},error) {
		return []byte(REFRESH_TOKEN_SECRET), nil
	})

	if err!=nil {
		return nil,err
	}
	 if _, ok:= result.Method.(*jwt.SigningMethodHMAC); !ok {
     return nil,err

	 }

	 if claims.ExpiresAt.Time.Before(time.Now()){
		return nil, errors.New("Token has expired")
	 }
	 return claims,nil
}