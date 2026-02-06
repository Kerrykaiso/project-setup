package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

var R *redis.Client
func ConnectRedis() {

if R !=nil{
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*1)
	defer cancel()
	if err:= R.Ping(ctx).Err(); err ==nil {
		return
	}
}


  rdb:=redis.NewClient(&redis.Options{
	Addr: "localhost:6379",

  })
  //defer rdb.Close()
ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
defer cancel()

    status:=rdb.Ping(ctx)
	if err:= status.Err(); err!=nil {
	log.Fatal("Error connecting to redis")
		return
	}
  R=rdb
 
 log.Println("Redis connected successfully")
}

func CloseRedis(){
	err:=R.Close()
	if err!=nil {
		log.Println("error closing redis")
	}
	log.Println("redis closed successfully")
}
var Channel *amqp.Channel
var Conn *amqp.Connection


//  func init(){
	
//  }

func ConnectRabbitMq() *amqp.Channel{
    RABBITMQ_URL,exist:=os.LookupEnv("RABBITMQ_URL")
	 if !exist {
		log.Fatal("No rabbitMq url")
	 }
   if Channel !=nil {
	 return Channel
   }
   conn, err:=amqp.Dial(RABBITMQ_URL)
   if err != nil {
	log.Fatal("Error connecting to rabbitMq")
   }
   Conn = conn
   ch, err:=conn.Channel()
   if err!=nil{
	log.Fatal("Error creating channel")
   }
   Channel=ch
   log.Println("RabbitMq connected")
   return  Channel
   
}

func CloseRabbitmq(){
	if Conn!=nil {
		Conn.Close()
	}
	if Channel!=nil {
		Channel.Close()
	}
	log.Println("Rabbitmq connection closed")
}