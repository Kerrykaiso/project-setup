package main

import (
	"log"
	"product-service/config"
	"product-service/events"
	"product-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main (){
	r:= gin.Default()
	if err:=godotenv.Load(".env"); err!=nil{
		log.Fatal("error loading env file")
	}
	config.ConnectRedis()
	config.ConnectDb()
	config.StartRabbitMq()
	defer config.CloseRabbitmq()
	events.ListenForEvents()
	config.DB.AutoMigrate(&config.ProductModel{})
    routes.ApiRoutes(r)
	log.Println("Product service running on port 8006")
	r.Run(":8006")
}