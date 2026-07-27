package config

import (
	"admin-service/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type AdminModel struct {
	UserId string  `json:"userId" gorm:"primaryKey"`
	Role string    `json:"role"`
	Password string `json:"password"`
	Email string    `json:"email" gorm:"uniqueIndex"`
}

func ConnectDb() {
	PORT := utils.GetEnv("PORT","5432")
	HOST := utils.GetEnv("HOST","localhost")
	DB_NAME := utils.GetEnv("DB_NAME","orderdb")
	DB_USER := utils.GetEnv("DB_USER","postgres")
	DB_PASSWORD := utils.GetEnv("DB_PASSWORD","kerryesua9@gmail.com")

    dbstring:= fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", HOST,DB_USER,DB_PASSWORD,DB_NAME,PORT)
    
		database , err :=gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		//log.SetFlags(log.Lshortfile)
		log.Fatal("error connecting to db")
	}

	database.AutoMigrate(&AdminModel{})
	DB=database
	log.Println("Database connected successfully")
}