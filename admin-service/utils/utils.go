package utils

import (
	"os"
	//jwt "github.com/golang-jwt/jwt/v4"
)


func GetEnv(key string, defaultValue string) string{
    if value,exist:= os.LookupEnv(key); exist{
       return value
	}
	return defaultValue
}



