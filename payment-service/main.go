package main

import (
	"log"
	"payment-service/config"
	"payment-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	if err:= godotenv.Load(".env"); err !=nil{
		log.Fatal("Error loading env file")
	}
	config.ConnectRedis()
	config.ConnectRabbitMq()
	defer config.CloseRabbitmq()
	defer config.CloseRedis()
	
	routes.PaymentRoutes(r)

	r.Run(":8009")
}