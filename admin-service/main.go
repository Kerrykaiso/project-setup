package main

import (
	"admin-service/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	r:=gin.Default()
	routes.ApiRoute(r)
    log.Println("Admin sever is running")
	r.Run(":8007")
}