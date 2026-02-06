package controllers

import (
	"bytes"
	_ "context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	_ "payment-service/config"
	"payment-service/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	//"
)

func PaymentHealth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Payment service up and running"})
}


func InitiatePaymentController (c *gin.Context){

  type paymentDetails struct{
		Email string `json:"email" validate:"required"`
		Amount int `json:"amount" validate:"required"`
	    ProductId string `json:"productId" validate:"required"`
	    DesignerName string `json:"designerName" validate:"required"`
		DesignerEmail string `json:"designerEmail" validate:"required"`
		ProductName string `json:"productName" validate:"required"`
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		Owner string `json:"owner" validate:"required"`
		Address string `json:"address" validate:"required"`
		State string `json:"state" validate:"required"`
	}
	type metadata struct{
		DesignerEmail string `json:"designerEmail"`
		PhoneNumber string   `json:"phoneNumber"`
		DesignerName string  `json:"designerName"`
		ProductName string   `json:"productName"`
		Owner string         `json:"owner"`
		Address string       `json:"address"`
	    ProductId string `json:"productId"`
		State string      `json:"state"`
	}

   type paystackBody struct{
		Email string `json:"email"`
		Amount int `json:"amount"`
		Metadata *metadata `json:"metadata"`
	}

	validate := validator.New()
    productData := &paymentDetails{}
	c.ShouldBindJSON(productData)
    if err:=validate.Struct(productData); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message": "Missing fields"})
		return
	}
    isLocked,err:=utils.IdempotencytLock(productData.ProductId)
	if !isLocked || err!=nil {
		c.JSON(http.StatusConflict,gin.H{"message":"This product is already being purchased"})
		return
	}
	payload := &paystackBody{
		Email: productData.Email,
		Amount: productData.Amount* 100,
		Metadata: &metadata{
			ProductName:productData.ProductName,
			ProductId: productData.ProductId,
			DesignerName: productData.DesignerName,
            DesignerEmail: productData.DesignerEmail,
			State: productData.State,
			Owner: productData.Owner,
			PhoneNumber:productData.PhoneNumber,
			Address: productData.Address,
		},
	}
        paystackPayload, err:=json.Marshal(payload)
		if err != nil {
			log.Println("Error marshalling payload")
			return
		}
	result, error:= http.NewRequest("POST","https://api.paystack.co/transaction/initialize",bytes.NewBuffer(paystackPayload))
    if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error making payment"})
		utils.IdempotencyUnlock(productData.ProductId)
		log.Println("Error initializing paystack payment")
		return
	}
	 PAYSTACK_API_KEY:=os.Getenv("PAYSTACK_SECRET_KEY")
	 BEARER_KEY := fmt.Sprintf("Bearer %v", PAYSTACK_API_KEY)

 	result.Header.Set("Authorization", BEARER_KEY)
	result.Header.Set("Content-Type", "application/json")
	client:= &http.Client{Timeout: 10*time.Second}
	response,err:=client.Do(result)

    if err!=nil {
		log.Println(err)
		utils.IdempotencyUnlock(productData.ProductId)
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Something went wrong"})
		return
	}
	defer response.Body.Close()

    body, err:=io.ReadAll(response.Body)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Something went wrong"})
		log.Println("Error reading paystack stream response")
		return
	}
	//send create order event

	
	bodyString := string(body)
	fmt.Println(bodyString)
	c.Data(response.StatusCode,"application/json", body)

}


type Paystack struct{
  Event string      `json:"event"`
  Data map[string]any  `json:"data"`
}

// type OrderStruct struct{
//  customerEmail string
//  ProductId string
//  OrderId string
//  CustomerName string
//  ProductName string
//  State string
//  PhoneNumber string
//  Address string
//  Reference string


// }
func PaystackWebhook(c *gin.Context){

	// type UpdateProduct struct{
	// 	Owner string `json:"owner"`
	// 	ProductId string `json:"productId"`
	// }
	// product := &UpdateProduct{}
    //  err:= c.ShouldBindJSON(product)
	//  if err !=nil {
	// 	log.Fatal(err)
	//  }
    //  body,err := json.Marshal(product)
	//  if err != nil {
	// 	log.Fatal(err)
	//  }
	// config.Channel.Publish(
	// 	"one-of-one-exchange",
	// 	"product.update",
	// 	false,
	// 	false,
    //    amqp.Publishing{
	// 	ContentType: "application/json",
	// 	Body: body,
	//    },

	// )
	// c.JSON(http.StatusOK,gin.H{"message":"payment published"})

	rawBody, err:=io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading request body"})
		return
	}
	signature:= c.GetHeader("X-Paystack-Signature")
	if signature=="" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "X-Paystack-Signature header missing"})
		return
	}

	PAYSTACK_SECRET_KEY :=os.Getenv("PAYSTACK_SECRET_KEY")

   if  !verifySignature(rawBody,signature,PAYSTACK_SECRET_KEY){
	  log.Println("Webhook received with INVALID signature!")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Signature verification failed"})
		return
   }
  var webhookData Paystack
  if err:=json.Unmarshal(rawBody,&webhookData);err !=nil{
	log.Printf("Error unmarshalling JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON payload"})
		return
  }

fmt.Println(webhookData)
c.JSON(http.StatusOK, gin.H{"message":"webhook received"})

    event:=webhookData.Event

	if event != "charge.success" {
		
	}

    orderId:= uuid.New().String()
    reference:=webhookData.Data["reference"].(string)
	amount := webhookData.Data["ëmail"].(string)
	customer := webhookData.Data["customer"].(map[string]any)
    email := customer["email"].(string)
	metadata,_ := webhookData.Data["metadata"].(map[string]any)
    productId:= metadata["productId"].(string)
//	designerName:= metadata["designerName"].(string)
	url:= utils.ProductUrl(productId)
	//designerEmail:= metadata["designerEmail"].(string)
	productName:= metadata["productName"].(string)
	owner:= metadata["owner"].(string)
	address:= metadata["address"].(string)
	state:= metadata["state"].(string)
	phoneNumber:= metadata["phoneNumber"].(string)

     orderDetails:= utils.CreateOrder{
		OrderId: orderId,
		Reference: reference,
		ProductId: productId,
		Url: url,
		ProductName: productName,
		CustomerName: owner,
		Address: address,
		State: state,
		Amount:amount,
		PhoneNumber: phoneNumber,
		CustomerEmail: email,
	 }
    update :=utils.ProductUpdate{
        ProductId: productId,
		Owner: owner,
	 }
	 err = utils.EmitCreateOrderEvent(orderDetails)
	 if err != nil {
		 log.Println("Error emitting create order event:", err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to emit create order event"})
		 return
	 }
	 err = utils.UpdateProduct(update)
	 if err != nil {
		 log.Println("Error emitting update product event:", err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to emit update product event"})
		 return
	 }
// 	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
// 	defer cancel()
// 	channel:=config.ConnectRabbitMq()
// 	error:=channel.PublishWithContext(ctx,"one-of-one-exchange","product.update",false,false,amqp091.Publishing{
// 		ContentType: "application/json",
// 		//Body: dataByte,
// 	})
// 	if error!=nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error":"error publishing"})
// 	    return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"error":"sent"})
// 	  log.Println("published!")
c.JSON(http.StatusOK, gin.H{"message":"webhook received"})
}


func verifySignature(body []byte, signature, secret string) bool {
	h:= hmac.New(sha512.New,[]byte(secret))

	h.Write(body)
    

	calculatedSignature:= hex.EncodeToString(h.Sum(nil))

	return  hmac.Equal([]byte(calculatedSignature),[]byte(signature))
}
