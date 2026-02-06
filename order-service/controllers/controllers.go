package controllers

import (
	"log"
	"net/http"
	"order-service/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderSericeHealthController(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{"message":"Order service running"})
}


func GetAllOrders(c *gin.Context){
	ctx:= c.Request.Context()
    no:= c.Query("no")

    number,err:= strconv.Atoi(no)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"message":"something went wrong","data":nil})
	  log.Println("error converting no to int")
	  return
	}
	var AllOrders []config.OrderModel

	query:= config.DB.WithContext(ctx).Where("status =?","pending").Order("created_at DESC").Limit(20)
	query= query.Where("no > ?", number)
	if err:=query.Find(&AllOrders).Error; err!=nil{
		if err == gorm.ErrRecordNotFound {
			 c.JSON(http.StatusBadRequest, gin.H{"message":"No orders found","data":nil})
	        log.Println("No orders yet")
	        return
		}
		 c.JSON(http.StatusInternalServerError, gin.H{"message":"something went wrong","data":nil})
	     log.Printf("%s", err.Error())
	    return
	}
   
	 lastNo:= len(AllOrders)-1
	 nextNo:=AllOrders[lastNo].No 
	c.JSON(http.StatusOK, gin.H{"message":"orders found","data":AllOrders,"NextNo":nextNo})
}


func GetOrderById (c *gin.Context){
  var order config.OrderModel

   orderId := c.Param("id")

  if err:=config.DB.Where("order_id=?",orderId).First(&order).Error; err!=nil{
    if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message":"Order not found","data":nil})
		log.Println("order not found")
		return
	}
	c.JSON(http.StatusInternalServerError,gin.H{"message":"something went wrong","data":nil})
	log.Printf("%s", err.Error())
	return
  }
  c.JSON(http.StatusFound, gin.H{"message":"order found","data":order})
}

func UpdateOrderStatus(c *gin.Context){
	var singleOrder config.OrderModel
	
     var update struct{
		Status *string `json:"status"`
	 }

	 orderId:= c.Param("orderId")
	 if err:=c.ShouldBindJSON(&update);err!=nil{
		log.Println("invalid user data")
		c.JSON(http.StatusBadRequest,gin.H{"message":"invalid or bad request"})
		return
	 }
	 updated:= map[string]interface{}{}
	 if update.Status!=nil {
		updated["status"] = *update.Status
	 }
	    result:=config.DB.Model(&singleOrder).Where("order_id=?",orderId).Updates(updated)
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Record not found")
			c.JSON(http.StatusNotFound,gin.H{"message":"Order Record not found"})
			return
		}
		if result.Error != nil {
			log.Printf("error: %s",result.Error.Error())
		 c.JSON(http.StatusInternalServerError,gin.H{"message":"something went wrong","data":nil})

		}
		c.JSON(http.StatusAccepted,gin.H{"message":"order status updated"})
	}

