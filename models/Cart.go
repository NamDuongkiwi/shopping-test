package models



type Cart struct{
	CartID	int `json:"cart_id"`
	ProductID	Product	`json:"product_id"`
	UserID		User	`json:"user_id"`
	CreatedAt	int64	`json:"created_at" gorm:"autoUpdateTime:milli"`
	Quantity	int		`json:"quantity"`
}