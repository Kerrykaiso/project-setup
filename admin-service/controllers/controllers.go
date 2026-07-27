package controllers

import (
	"admin-service/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-redis/redis/v8"
	_ "github.com/google/uuid"
	_ "github.com/rabbitmq/amqp091-go"
)
 type Register struct{
   Email string `json:"email" binding:"required,email,min=5,max=20"`
   Role string `json:"role" binding:"required,max=8"`
   Password string `json:"password" binding:"required,min=5,max=20"`
 }
func HealthController(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message":"Admin server up and running"})
}


func RegisterAdminController(c *gin.Context){
  var register Register

  if err:=c.ShouldBindJSON(&register); err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"message":"Incorrect input",
		"error" : err.Error(),
})
  log.Println(err.Error())
	return
  }
  
  foundUser:=config.AdminModel{}

  if err:=config.DB.Where("email =?",register.Email).First(&foundUser).Error; err==nil{
    log.Println("Email already in use")
	c.JSON(http.StatusForbidden,gin.H{"message":"Email already in use"})
	return
  }
  

  c.JSON(http.StatusCreated,gin.H{"message":"admin registered"})
}