package controllers

import (
	"context"
	"designer-service/config"
	"designer-service/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
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
	"userId":newDesigner.UserId,
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
 fmt.Println(data)
    if err:= config.DB.Where("email = ?", data.Email).First(foundUser).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
	   return
	}

  if err:= bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(data.Password)); err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error":"Incorrect email or password"})
	return
  }
   accessToken,refreshToken,err := utils.GenerateAcessAndRefreshToken(foundUser.UserId, foundUser.Email,foundUser.Name )
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


func GetInventory(c *gin.Context){
  var inventories []config.Inventory

  if err:= config.DB.Find(&inventories).Error; err != nil{
    c.JSON(http.StatusNotFound, gin.H{"error":"Inventory empty"})
	return
  }
  c.JSON(http.StatusAccepted,gin.H{"message":inventories})
}


func AddInventory (c *gin.Context){
  var inventory []config.Inventory
 if err:= c.ShouldBindJSON(&inventory); err!=nil{
	c.JSON(http.StatusConflict, gin.H{"message": "something went wrong"})
	return
 }
 if err:=config.DB.Create(&inventory).Error; err !=nil{
   	c.JSON(http.StatusConflict, gin.H{"message": "something went wrong"})
    return
 }
  c.JSON(http.StatusCreated, gin.H{"message": "New blanks added"})

}


func SubmitDesign(c *gin.Context){
	val,exist:=c.Get("user")
	if !exist {
		 c.JSON(http.StatusConflict, gin.H{"message": "User not logged in"})
		 return
	}
	user,ok:=val.(*utils.LoginStruct)
	if !ok {
		fmt.Println("invalid login type")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid auth details"})
		return
	}
   fmt.Println("designer_email",user.Email)
	ctx,cancel:=context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
    var submissionDTO config.SubmissionDTO
	err:= c.ShouldBindJSON(&submissionDTO)
	if err !=nil{
		fmt.Println("err",err)
	    c.JSON(http.StatusConflict, gin.H{"message": "Missing form field(s)"})
	    return
	}
  designs,_:= json.Marshal(submissionDTO.Designs)
 
  
   txErr:=config.DB.Transaction(func (tx *gorm.DB) error{
   
	submission:= config.Submission{
		Color: submissionDTO.Color,
		Size: submissionDTO.Size,
		DesignerName: submissionDTO.DesignerName,
		FrontImage: submissionDTO.FrontImage,
		BackImage: submissionDTO.BackImage,
		Designs: datatypes.JSON(designs),
		DesignerEmail: user.Email,
	}
	 if err:= tx.WithContext(ctx).Create(&submission).Error; err!=nil{
	  return err
	 }

	  var inventory config.Inventory
	  result:=tx.WithContext(ctx).Model(&inventory).Where("color =? AND size =? AND quantity > 0",
	   submissionDTO.Color,submissionDTO.Size).Update("quantity",gorm.Expr("quantity-?",1))
	  
      if result.Error != nil {
      return result.Error 
      }

      if result.RowsAffected == 0 {
        return fmt.Errorf("item is out of stock or combination does not exist")
       }
	  return nil
   })

   if txErr!=nil {
	c.JSON(http.StatusBadRequest, gin.H{"message": "submission failed"})
	log.Println(txErr.Error())
	return
   }
    c.JSON(http.StatusCreated, gin.H{"message": "Design submitted!"})
}


func UploadDesign(c *gin.Context){
details:= utils.GenerateSignature()
c.JSON(http.StatusAccepted,gin.H{"message":details})
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
 // complete

}

func Logout(c *gin.Context){
	c.SetCookie("accessToken", "", -1,"","",false,false)
	c.SetCookie("refreshToken", "", -1,"/api/refresh","",false,false)
	c.JSON(http.StatusOK, gin.H{"Message":"Logout successful"})
}

