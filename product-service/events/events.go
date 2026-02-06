package events

import (
	"encoding/json"
	"fmt"
	"log"
	"product-service/config"
	"product-service/service"
)



func ListenForEvents(){
   ch,err:=config.StartRabbitMq()
   if err!=nil {
	log.Fatal("Could not start rabbitMq")
   }
     exchangeError:=ch.ExchangeDeclare(
	    "one-of-one-exchange",
         "direct",
		 true,
		 false,
		 false,
		 false,
		 nil,
	)
	if exchangeError !=nil{
      log.Println("error creating exchange")
	}

	q,err:=ch.QueueDeclare(
		"product-updated",
		true,
		false,
		false,
		true,
		nil,
	)

	if err != nil {
       log.Println("error creating update product queue")
	}
 
	bindError:=ch.QueueBind(
		q.Name,
		"product.update",
		"one-of-one-exchange",
		false,
		nil,
	)
	if bindError!=nil {
		log.Println("error binding update product queue")
	}
	msgs,err:=ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
	log.Println("error consuming message event")
	}

	go func() {
		log.Println(" product event consumer listening for events")
		for msg := range msgs{
			log.Println("i ran")
			//log.Println(msg.Body)
			var data service.Data
			err:=json.Unmarshal(msg.Body, &data)
			if err != nil {
				msg.Nack(false,false)
				continue
			}
			errr:=service.UpdateProductInDB(data)
			if errr!=nil {
				fmt.Println("error ooo!")
				fmt.Print(errr)
				msg.Nack(false,true)
				continue
			}
			log.Println("Product updated")
			msg.Ack(false)
			fmt.Println(data)
		}
	}()
}