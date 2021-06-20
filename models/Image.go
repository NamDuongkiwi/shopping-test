package models

type Image struct{
	ProductId	Product `json:"product_id"`
	ImageUrl	string `json:"url"`
}