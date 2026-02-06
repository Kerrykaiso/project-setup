package main

import (
	"log"
	"order-service/config"
	"order-service/events"

	"order-service/queues"
	"order-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err:=godotenv.Load(".env")
	if err!=nil {
		log.Fatal("failed to load env files")
	}
  config.ConnectDb()	
  config.StartRabbitMq()
  config.ConnectRedis()
  defer config.R.Close()
  go queues.StartJobQueue()
  defer config.CloseRabbitmq()
 go events.ListenForEvents()
  r:= gin.Default()
  routes.ApiRoutes(r)
  log.Println("Order service running")
  r.Run(":8008")
}