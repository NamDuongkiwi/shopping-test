package models

type Review struct{
	ReviewID int `json:"review_id"`
	ProductID Product `josn:"product_id"`
}