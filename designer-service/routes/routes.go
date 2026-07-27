package routes

import (
	"designer-service/controllers"
	"designer-service/middleware"

	"github.com/gin-gonic/gin"
)


func ApiRoutes(r *gin.Engine) {
	r.GET("/api/designer-service-health",controllers.HealthController)
	r.POST("/api/signup",controllers.SignUp)
	r.POST("/api/login",controllers.Login)
	r.POST("/api/add-inventory", controllers.AddInventory)
	r.GET("/api/get-inventory", controllers.GetInventory)
	r.GET("/api/get-signature",middleware.AuthMiddleware(),controllers.UploadDesign)
	r.POST("/api/create-submission",middleware.AuthMiddleware(),controllers.SubmitDesign)
	r.GET("/api/refresh")
}