package routes

import (
	"designer-service/controllers"

	"github.com/gin-gonic/gin"
)


func ApiRoutes(r *gin.Engine) {
	r.GET("/api/designer-service-health",controllers.HealthController)
	r.POST("/api/signup",controllers.SignUp)
	r.POST("/api/login",controllers.Login)
	r.GET("/api/refresh")
}