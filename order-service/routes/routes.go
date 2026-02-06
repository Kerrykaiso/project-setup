package routes

import (
	"order-service/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/order-service-health",controllers.OrderSericeHealthController)
	r.GET("/allOrders", controllers.GetAllOrders)
	r.GET("/order/:id", controllers.GetOrderById)
}