package config

import (
	"context"
	"fmt"
	"log"
	"order-service/utils"
	"time"

	"github.com/go-redis/redis/v8"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
type OrderModel struct{
  ProductId string `json:"productId" gorm:"not null"`
  Amount string `json:"amount" gorm:"not null"`
  No int `json:"no" gorm:"autoIncrement"`
  Country string `json:"country" gorm:"default:'Nigeria'"`
  OrderId string `json:"orderId" gorm:"uniqueIndex;not null;primaryKey"`
  CustomerName string `json:"customerName" gorm:"not null"`
  ProductName string `json:"productName" gorm:"not null"`
  State string `json:"state" gorm:"not null"`
  PhoneNumber string `json:"phoneNumber" gorm:"not null"`
  Address string `json:"location" gorm:"not null"`
  CustomerEmail string `json:"customerEmail"`
  Reference string `json:"reference" gorm:"not null"`
  Url string `json:"url" gorm:"not null"`
  Status string `json:"status" gorm:"not null;default:pending"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

func ConnectDb (){
	PORT := utils.GetEnv("PORT","5432")
	HOST := utils.GetEnv("HOST","localhost")
	DB_NAME := utils.GetEnv("DB_NAME","orderdb")
	DB_USER := utils.GetEnv("DB_USER","postgres")
	DB_PASSWORD := utils.GetEnv("DB_PASSWORD","kerryesua9@gmail.com")
	dbstring:= fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", HOST,DB_USER,DB_PASSWORD,DB_NAME,PORT)
	database , err :=gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		//log.SetFlags(log.Lshortfile)
		log.Fatal("error conneting to db")
	}

	database.AutoMigrate(&OrderModel{})
	DB=database
	log.Println("Database connected successfully")
}


var Channel *amqp.Channel
var Conn *amqp.Connection
func StartRabbitMq()(*amqp.Channel, error){
 if Conn==nil {
	   conn, err:=amqp.Dial(utils.GetEnv("RABBITMQ_URL","amqp://localhost:5672"))
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


var R *redis.Client

func ConnectRedis(){
  if R !=nil{
   ctx,cancel:= context.WithTimeout(context.Background(),time.Second*1)
   defer cancel()
   if err:=R.Ping(ctx).Err(); err==nil{
     return
   }
  }

    rdb:=redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
  })

ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
defer cancel()

    _,err:=rdb.Ping(ctx).Result()
	if  err!=nil {
	log.Fatal("Error connecting to redis")
		return
	}
  R=rdb
 fmt.Println("Redis connected successfully")
}