package routes

import (
	"payment-service/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	// Define payment-related routes here
	r.GET("/payment-service-health",controllers.PaymentHealth)
	r.POST("/initialize-payment", controllers.InitiatePaymentController)
	r.POST("/payment-webhook",controllers.PaystackWebhook)
}