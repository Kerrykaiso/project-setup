package service

import (
	"errors"
	"log"
	"product-service/config"

	"github.com/go-playground/validator/v10"
)

type Data struct {
	Owner string `json:"owner" validate:"required"`
	Id string `json:"productId" validate:"required"`
}



func UpdateProductInDB( data Data) error {

	//validate data
	  validate := validator.New()

    if err:= validate.Struct(data); err!=nil{
		log.Printf("validation error %v", err)
       return errors.New("something went wrong")
	}

updates:=map[string]interface{}{
	"owner": data.Owner,
	"status": "sold",
}
 result:= config.DB.Model(&config.ProductModel{}).Where("product_id=?", data.Id).Updates(updates)
	
 if result.Error != nil{
	log.Println("Failed to update product")
	return errors.New("failed to update product")
 }

 if result.RowsAffected == 0 {
	log.Println("could not find product")
	return errors.New("failed to find product")
 }
 return nil
}