package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/datatypes"
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

type Inventory struct {
	Color string `json:"color"`
	Size string `json:"size"`
	Quantity int `json:"quantity" gorm:"check:quantity>=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Submission struct {
	Color string `json:"color"`
	Size string `json:"size"`
	DesignerName string `json:"designerName"`
	DesignerEmail string `json:"designerEmail"`
	Reserved bool  `json:"reserved"`
	FrontImage string `json:"frontImage"`
	BackImage string `json:"backImage"`
	Designs datatypes.JSON `json:"designs"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}
type SubmissionDTO struct {
 	Color string `json:"color" binding:"required"`
	Size string `json:"size" binding:"required"`
	DesignerName string `json:"designerName" binding:"required"`
	Reserved bool  `json:"reserved"`
	FrontImage string `json:"frontImage" binding:"required"`
	BackImage string `json:"backImage" binding:"required"`
	Designs []string `json:"designs" binding:"required"`
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

