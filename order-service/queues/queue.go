package queues

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"order-service/config"
	"order-service/events"

	//"order-service/service"
	"order-service/structs"
)

func StartJobQueue(){

	ctx:= context.Background()
	log.Println("create order job queue started")
	for {
       result,err:= config.R.BRPop(ctx,0,"order_queue").Result()
	   if err != nil {
		log.Println(err)
		continue
	   }
	   data:= result[1]

	  var job events.Job
	  err=json.Unmarshal([]byte(data),&job)

	  if err !=nil {
		log.Println("error unmarshalling job data",err)
		continue
	  }
	   fmt.Println("create order data",string(job.Payload))
	//    if err:=ProcessJob(job);err!=nil{
	// 	log.Println("failed to process order job")
	// 	continue
	//    }
        
	}
}

func ProcessJob(data  events.Job)error{
	var order structs.CreateOrder
  err:=json.Unmarshal(data.Payload,&order)
  if err != nil {
	return err
  }
  //err = service.CreateOrderService(order)
  fmt.Println(order)
  return err
}