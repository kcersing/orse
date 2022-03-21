package main

import (
	"orse/common/models"
	"orse/internal/database"
)

func main() {
	client := database.Open()
	client.AutoMigrate(&models.Menu{})
	client.AutoMigrate(&models.User{}, &models.UserDetail{})
	client.AutoMigrate(&models.ProductCate{}, &models.Product{}, &models.ProductSpecs{}, models.ProductSpecsItem{},
		&models.ProductSpecsItemValues{}, &models.ProductAttributeKey{}, &models.ProductAttributeValue{})
	client.AutoMigrate(&models.Order{}, &models.OrderAmounts{}, &models.OrderSetting{}, models.OrderPay{}, &models.OrderItem{})

}
