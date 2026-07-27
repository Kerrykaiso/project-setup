package routes

import (
	"admin-service/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(r *gin.Engine) {
   r.GET("/api/admin-service-health",controllers.HealthController)
   r.POST("/api/register-admin",controllers.RegisterAdminController)
}