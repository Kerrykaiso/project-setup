package events

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"order-service/config"
)

type Job struct{
     Payload  json.RawMessage `json:"payload"`
     Attempts int  `json:"attempts"`
     MaxAttempts int `json:"maxAttempts"`

}
func ListenForEvents(){
	fmt.Println("Listening for order events")
	err:=config.Channel.ExchangeDeclare(
		"one-of-one-exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declaring order exchange:%v",err)
	}

	q,err:=config.Channel.QueueDeclare(
		"order-queue",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
	 log.Fatalf("Error declaring order queue:%v",err)
	}

	bindError:=config.Channel.QueueBind(
		q.Name,
		"create.order",
		"one-of-one-exchange",
		false,
		nil,
	)
	if bindError!=nil {
		log.Println("error binding update product queue")
	}

	
	
   msgs,err:=config.Channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
	  log.Println("error consuming create order messages")
	}
   
	for msg:= range msgs {

		job:= Job{
			Payload: json.RawMessage(msg.Body),
			Attempts: 1,
			MaxAttempts: 5,
		}

	    newJob,errr:=json.Marshal(job)
		 if errr != nil {
			log.Println("error marshalling job",errr)
			continue
		 }
	
		err=config.R.LPush(context.Background(),"order_queue",newJob).Err()
		if err != nil {
			msg.Nack(false,true)
		}

		msg.Ack(false)
	}
	
}

