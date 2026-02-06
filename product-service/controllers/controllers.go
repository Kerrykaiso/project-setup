package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-service/config"
	"product-service/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ProductHealth(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{"message":"Product service up and running"})
}

func GetAllProducts (c *gin.Context){
   //var pageLimit = 10
   var cursor =c.Query("cursor")
   var page = c.Query("page")
   var color = c.Query("color")
   var createdAt = c.Query("createdAt")
   var cacheKey = fmt.Sprintf("products:page:%v:color:%v", page,color)
   if cursor=="" {
	 cursor ="start"
   }
 
   //check cache
    
      ctx:=c.Request.Context()
	  redisCtx, cancel:= context.WithTimeout(ctx, time.Second*3)
	  defer cancel()
         if cached,err:= config.R.Get(redisCtx,cacheKey).Result(); err==nil{
			fmt.Printf("cache hit for page %v/n",page)
			var products []config.ProductModel
			if err:= json.Unmarshal([]byte(cached), &products); err==nil{
				c.JSON(http.StatusOK,gin.H{"cache":true,"data":products})
				return
			}
		 }

		 //get products from db
       
       var products []config.ProductModel
	   query:= config.DB.WithContext(ctx).Where("status=?","available").Order("cursor DESC, created_at DESC").Limit(5)
        
	   if color !="" {
		query = query.Where("color=?", color)
	   }
	   if cursor !="start" && createdAt!=""{
		 t,err:=time.Parse(time.RFC3339,createdAt)
		 if err != nil {
			log.Println("Error parsing createdAt")
			c.JSON(http.StatusBadRequest,gin.H{"cache":false,"message":"bad request","data":products})
			return 
		 }
		 num, err:=strconv.Atoi(cursor)
		 if err != nil {
			log.Println("error converting cursor to int")
			c.JSON(http.StatusBadRequest,gin.H{"cache":false,"message":"something went wrong"})
			return
		 }
		query = query.Where("(created_at < ?) OR (created_at =? AND cursor < ?)",t, t, num)
	   }
     if err:=query.Find(&products).Error; err !=nil{
       c.JSON(http.StatusBadRequest, gin.H{"cache":false, "data":products, "message":"products not found"})
	   return
	 }
	 
	 // save to cache
	 product, err:= json.Marshal(products)
	 if err!=nil {
		 log.Println("Error saving products to cache")
		}
		config.R.Set(ctx,cacheKey,product, 1*time.Minute )
		log.Printf("%v, has been saved to cache", cacheKey)

     prodLen:=len(products)
	 var created string
	 latest:=products[prodLen-1]
	 next_cursor := latest.Cursor
	 //created = latest.CreatedAt.Format(time.RFC3339)
     created = latest.CreatedAt.UTC().Format(time.RFC3339)
	c.JSON(http.StatusOK, gin.H{"cache":false, 
	"data":products,
	 "message":"products found",
	 "createdAt":created,
	 "cursor":next_cursor,
	  
	})
	}


func AddProducts (c *gin.Context){

  type productValue struct {
  ProductName string `json:"productName" validate:"required"`
  ProductId string `json:"productId"`
  DesignerEmail string `json:"designerEmail" validate:"required"`
  DesignerName string `json:"designerName" validate:"required"`
  Status string `json:"status"`
  Size string `json:"size" validate:"required"`
  Color string `json:"color" validate:"required" `
  Cursor uint `json:"cursor"`
  Front string `json:"frontImage" validate:"required"`
  Owner string `json:"owner"`
  Back string `json:"backImage" validate:"required"`
  }
     
 product:= &productValue{}
 
 var validate = validator.New()
     err:= c.ShouldBindJSON(product)

  if err := validate.Struct(product); err!= nil{
        fmt.Println("Validation error")
		var errMessage []string
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		for _,verr:=range err.(validator.ValidationErrors){
          errMessage=append(errMessage,fmt.Sprintf("error on field %s: %s\n",verr.Field(),verr.Tag())) 
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors":errMessage })
		return
  }
	 if err != nil {
		fmt.Println("Missing field")
	 }

	 productId := uuid.New().String()

	 productData := &config.ProductModel{
		ProductName: product.ProductName,
		ProductId: productId,
		DesignerEmail: product.DesignerEmail,
		DesignerName: product.DesignerName,
		Back: product.Back,
		Front: product.Front,
		Size: product.Size,
		Status: product.Status,
		Owner: product.Owner,
		Color: product.Color,

	 }
	 var existingProdut config.ProductModel
	 if err:= config.DB.Where("product_name=?",product.ProductName).First(&existingProdut).Error; err == nil{
       log.Println("product Name already in use")
	  c.JSON(http.StatusAccepted, gin.H{"error":"Product name already in use"})
      return
	 }
  if err:=config.DB.Create(productData).Error; err !=nil{
     fmt.Println("Error creating product")
	 c.JSON(http.StatusAccepted, gin.H{"error":err.Error()})
	 return
  }

  c.JSON(http.StatusAccepted, gin.H{"message":"product created"})
}


func GetProductByColor(c *gin.Context){

	  var products []config.ProductModel
	  ctx:=c.Request.Context()
	  context, cancel:= context.WithTimeout(ctx, time.Second*3)
	  defer cancel()
     color := c.Query("color")
     query:=config.DB.WithContext(ctx).Where("status=?", "available").Order("created_at DESC")
     if color =="" {
	 if err := query.Find(&products).Error; err!=nil {
      fmt.Println("Error fetching products")
	  return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Products fetched", "data":products})
  }
  cacheKey := fmt.Sprintf("product:color:%v",color)

 if data,err:=config.R.Get(context,cacheKey).Result(); err==nil{
    if err:=json.Unmarshal([]byte(data), &products); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Something went wrong"})
		log.Println("Error unmarshalling json")
		return
	}
    c.JSON(http.StatusOK,gin.H{"cache":true,"message":"Fetched product with colors", "data":products})
    return
 }
 //cache miss

  query= config.DB.WithContext(ctx).Where("color=?", color).Order("created_at DESC")
  if err:= query.Find(&products).Error; err!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find color","data":products})
	return
  }
   // save to cache
	 product, err:= json.Marshal(products)
	 if err!=nil {
		 log.Println("Error saving products to cache")
		}
		config.R.Set(ctx,cacheKey,product, 1*time.Minute )
		log.Printf("%v, has been saved to cache", cacheKey)

    c.JSON(http.StatusOK, gin.H{"message": "Colors fetched successfully","data":products})

}


func GetProductById (c *gin.Context){
	var singleProduct config.ProductModel
	ctx:= c.Request.Context()
  productId := c.Param("productId")
  
    if err:=config.DB.WithContext(ctx).Where("product_id=?",productId).First(&singleProduct).Error; err!=nil{
		log.Println("Error fetching product by Id")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch product","data":nil})
		return
	}
      c.JSON(http.StatusOK, gin.H{"message": "Product fetched successfully","data":singleProduct})

}


func UpdateProductController (c *gin.Context){

 data := &service.Data{}
 c.ShouldBindJSON(data)

 if err:= service.UpdateProductInDB(*data); err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	return
 }
 c.JSON(http.StatusOK, gin.H{"message":"Product updated successfully"})
}
