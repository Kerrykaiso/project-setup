package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


type UserModel struct {
	UserId string `json:"userId" gorm:"primaryKey"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email" gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConnectDb(){
	dsn := "host=localhost user=postgres password=kerryesua9@gmail.com dbname=designerdb port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Problem connecting to database", err)
	}
	DB = database
   fmt.Println("Database connected successfully")
} 