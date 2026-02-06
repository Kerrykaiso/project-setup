package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var R *redis.Client
var DB *gorm.DB


type ProductModel struct {
  ProductName string `json:"productName" gorm:"uniqueIndex"`
  ProductId string `json:"productId" gorm:"primaryKey"`
  DesignerEmail string `json:"designerEmail"`
  DesignerName string `json:"designerName"`
  Status string `json:"status" gorm:"default:'available'"`
  Size string `json:"size" gorm:"default:XL"`
  Color string `json:"color"`
  Cursor uint `json:"cursor" gorm:"uniqueIndex;autoIncrement"`
  Front string `json:"frontImage"`
  Owner string `json:"owner"`
  Back string `json:"backImage"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

func ConnectDb(){
	DB_NAME := os.Getenv("DB_NAME")
	//dsn:="host=localhost user=postgres password=kerryesua9@gmail.com dbname=productdb port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=localhost user=postgres password=kerryesua9@gmail.com dbname=%v port=5432 sslmode=disable", DB_NAME)
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Problem connecting to database", err)
	}
	postgres,err:= database.DB()
	if err != nil {
		log.Fatal("error configuring db")
	}
	postgres.SetMaxIdleConns(7)
	postgres.SetMaxOpenConns(15)
	postgres.SetConnMaxLifetime(10*time.Minute)
	
	DB = database
   log.Println("Database connected successfully")
}

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

ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
defer cancel()

    status:=rdb.Ping(ctx)
	if err:= status.Err(); err!=nil {
	log.Fatal("Error connecting to redis")
		return
	}
  R=rdb
 fmt.Println("Redis connected successfully")
}



var Channel *amqp.Channel
var Conn *amqp.Connection
func StartRabbitMq()(*amqp.Channel, error){
 if Conn==nil {
	   conn, err:=amqp.Dial(os.Getenv("RABBITMQ_URL"))
	   if err != nil {
		log.Fatal("Error connecting to rabbitMq")
	   }
	   Conn=conn
 }
   if Channel !=nil {
	 return Channel, nil
   }
   channel, err:=Conn.Channel()
   if err!=nil{
	log.Fatal("Error creating channel")
   }
   Channel=channel
   log.Println("RabbitMq connected successfully")
   return  Channel, nil
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