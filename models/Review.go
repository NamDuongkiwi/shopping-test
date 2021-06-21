package models

type Review struct{
	ReviewID 	int	`json:"review_id"`
	ProductID 	int `json:"product_id"`
	Review		string `json:"reivew"`
	Rating		float32 `json:"rating"`	
}