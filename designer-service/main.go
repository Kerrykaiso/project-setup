package main

import (
	"designer-service/config"
	"designer-service/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	err:=godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
 r:= gin.Default()
 proxyerr := r.SetTrustedProxies([]string{
    "127.0.0.1",
    "::1",
})
if proxyerr != nil {
    log.Fatal(err)
}
  config.ConnectDb()
  config.DB.AutoMigrate(&config.UserModel{},&config.Inventory{},&config.Submission{})
   routes.ApiRoutes(r)
 r.Run(":8005")
}