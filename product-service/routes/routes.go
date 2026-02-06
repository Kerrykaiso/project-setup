package routes

import (
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
  r.GET("/product-service-health", controllers.ProductHealth)
  r.GET("/all-products", controllers.GetAllProducts)
  r.GET("/product-color", controllers.GetProductByColor)
  r.GET("/product/:productId", controllers.GetProductById)
  r.POST("/products", controllers.AddProducts)
}