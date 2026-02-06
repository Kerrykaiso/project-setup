package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"payment-service/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


var BASE_URL = "http://oneofone.com"

type CreateOrder struct{
  ProductId string
  PhoneNumber string
  OrderId string
  CustomerName string
  ProductName string
  State string
  Address string
  CustomerEmail string
  Amount string
  Reference string
  Url string
}
type ProductUpdate struct{
	Owner string
	ProductId string

}

func GetEnv(key string, defaultValue string) string{
    if value,exist:= os.LookupEnv(key); exist{
       return value
	}
	return defaultValue
}

func IdempotencytLock (key string)(bool,error){
	fmt.Println("Acquiring Idempotency Lock")
	ctx:= context.Background()
	result:=config.R.SetNX(ctx,key,"1",time.Duration(2*time.Minute))
	ok,err:= result.Result()
	if err != nil {
		fmt.Print(err)
		return false,errors.New("error acquiring idempotency lock")
	} 
	if !ok {
		return false,errors.New("idempotencty Key already locked")
	}
   return true, nil
}

func IdempotencyUnlock(key string) error {
    ctx := context.Background()
    err := config.R.Del(ctx, key).Err()
    if err != nil {
        fmt.Printf("Failed to unlock key %s: %v\n", key, err)
        return err
    }
    fmt.Println("Idempotency Lock released for:", key)
    return nil
}

func ProductUrl(productId string) string {
 url:= fmt.Sprintf("%s/%s",BASE_URL,productId)
 return url
}

func EmitCreateOrderEvent(orderDetails CreateOrder)error{
	orderBody,err:=json.Marshal(orderDetails)
	if err != nil {
		log.Println("error marshalling orders")
		return err
	}
  err = config.Channel.Publish(
		"one-of-one-exchange",
		"create.order",
		false,
		false,
       amqp.Publishing{
		ContentType: "application/json",
		Body: orderBody,
	   },
  )
 log.Println("sending order create event")
  return err
}


func UpdateProduct(product ProductUpdate)error{
	updateProductBody,err := json.Marshal(product)
	if err != nil {
		log.Printf("error updating product %v",err)
		return err
	}
	err = config.Channel.Publish(
		"one-of-one-exchange",
		"product.update",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: updateProductBody,
		},
	)
	fmt.Println("emitting update product!!")
	return err
}