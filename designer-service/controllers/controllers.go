package controllers

import (
	"designer-service/config"
	"designer-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type DesignerData struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
	Email string `json:"email" validate:"required,email"`
}
func HealthController(c *gin.Context) {
	
 c.JSON(200, gin.H{"message":"designer server up and running"})
}

var validate = validator.New()

func SignUp(c *gin.Context){
  data := &DesignerData{}

  c.ShouldBindJSON(data)
  
  err:= validate.Struct(data)

 if err!=nil {
	c.JSON(401, gin.H{"error": err.Error()})
	return
 }

 //check if email exists 
 var existingUser config.UserModel 

  if err:= config.DB.Where("email =?", data.Email).First(&existingUser).Error; err == nil {
   fmt.Println("This email is already in use")
   c.JSON(401, gin.H{"error": "This email is already in use"})
   return
  }
  hassPassword,err := utils.Hashpassword(data.Password)
  if err != nil {
	fmt.Println(err.Error())
	c.JSON(400, gin.H{"error": "Something went wrong"})
  }
  fmt.Println(hassPassword)
  userId := uuid.New().String() 

  newDesigner := config.UserModel{
	UserId: userId,
	Email: data.Email,
	Password: hassPassword,
	Name: data.Name,
  }
 if err := config.DB.Create(&newDesigner).Error; err!=nil{
  c.JSON(http.StatusInternalServerError, gin.H{"message":"error creating user","details":err.Error()})
  return
 }
   c.JSON(http.StatusAccepted,gin.H{"message":"Designer created successfully","data": gin.H{
	"useId":newDesigner.UserId,
	"email":newDesigner.Email,
	"name":newDesigner.Name,
   }})
}



func Login(c *gin.Context){
  foundUser := &config.UserModel{}

  
  type LoginData struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
  }
   var data LoginData
    c.ShouldBindJSON(&data)

	  err:= validate.Struct(data)

 if err!=nil {
	c.JSON(401, gin.H{"error": "Invalid input"})
	return
 }
    if err:= config.DB.Where("email = ?", data.Email).First(foundUser).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
	   return
	}

  if err:= bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(data.Password)); err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error":"Incorrect email or password"})
	return
  }
   accessToken,refreshToken,err := utils.GenerateAcessAndRefreshToken(foundUser.Name, foundUser.UserId, foundUser.Email)
   if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error":"Error generating access token"})
	return
   }

   c.SetCookie(
	"accessToken",
	accessToken,
	3600*24,
	"/",
	"",
    false,
	true,
   )
    c.SetCookie(
	"refreshToken",
	refreshToken,
	3600*168,
	"/api/refresh",
	"",
    false,
	true,
   )
	c.JSON(http.StatusOK, gin.H{"access": accessToken,"refresh":refreshToken})
}



func Refresh(c *gin.Context){
 cookie,err:= c.Cookie("refreshToken")
 if err !=nil {
	c.JSON(http.StatusForbidden, gin.H{"error":err.Error()})
	return
 }
 if cookie=="" {
	c.JSON(http.StatusForbidden, gin.H{"error":"Missing refresh token"})
	return
 }

}

func Logout(c *gin.Context){
	c.SetCookie("accessToken", "", -1,"","",false,false)
	c.SetCookie("refreshToken", "", -1,"/api/refresh","",false,false)
	c.JSON(http.StatusOK, gin.H{"Message":"Logout successful"})
}