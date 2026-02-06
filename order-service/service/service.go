package service

import (
	"log"
	"order-service/config"
	"order-service/structs"
)


func CreateOrderService(order structs.CreateOrder) error {
 newOrder := &config.OrderModel{
	ProductId: order.ProductId,
	OrderId: order.OrderId,
	Reference: order.Reference,
	ProductName: order.ProductName,
	PhoneNumber: order.PhoneNumber,
	State: order.State,
	CustomerEmail: order.CustomerEmail,
	Address: order.Address,
	Url: order.Url,
	CustomerName: order.CustomerName,}

	result := config.DB.Create(newOrder)
	if result.Error ==nil {
		log.Println("Order created")
	}
    return result.Error
}