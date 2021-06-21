package models

type PriceUpdade struct{
	ProductID 	int 	`json:"product_id"`
	CreatedAt 	int64 	`json:"created-at" gorm:"autoUpdateTime:milli"`
	Price		int64	`json:"price"`
}