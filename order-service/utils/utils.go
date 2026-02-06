package utils

import (
	"os"
)


func GetEnv(key string, defaultValue string) string{
    if value,exist:= os.LookupEnv(key); exist{
       return value
	}
	return defaultValue
}




