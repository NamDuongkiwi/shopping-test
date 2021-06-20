package models


type Product struct {
	ID          int    		`json:"id"`
	Category_id Category    `json:"category_id" gorm:"foreignKey:CategoryRefer"`
	Name        string 		`json:"name"`
	Price       int64  		`json:"price"`
	IsSale      bool   		`json:"is_sale"`
	CreatedAt   int64  		`json:"created-at" gorm:"autoUpdateTime:milli"`
	Image		[]Image 	`json:"images"`
	Rating		float32		`json:"rating"`
}