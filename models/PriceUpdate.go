package models

type PriceUpdade struct{
	ProductID 	int 	`json:"product_id"`
	UpdatedAt 	string 	`json:"update-at"`
	Price		int64	`json:"price"`
}