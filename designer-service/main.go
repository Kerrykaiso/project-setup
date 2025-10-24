package main

import (
	"designer-service/config"
	"designer-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	err:=godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
 r:= gin.Default()
  config.ConnectDb()
  config.DB.AutoMigrate(&config.UserModel{})
   routes.ApiRoutes(r)
 r.Run(":8005")
}