package middleware

import (
	"designer-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware()gin.HandlerFunc {
   return func (c *gin.Context){
	accessToken,tokenError:= c.Cookie("accessToken")
	 if tokenError!=nil {
	   c.JSON(http.StatusConflict, gin.H{"message": "User not logged in"})
	   c.Abort()
	   return
	 }
    user,verificationError:= utils.VerifyToken(accessToken)
	if verificationError!=nil {
		fmt.Println(verificationError)
		c.JSON(http.StatusConflict, gin.H{"message": "Invalid or expired token"})
		c.Abort()
        return
	}
	c.Set("user",user)
	c.Next()
   }
}