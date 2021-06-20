package models


type Product struct {
	ID          int    		`json:"id"`
	Category_id int		    `json:"category_id" gorm:"foreignKey:CategoryRefer"`
	Name        string 		`json:"name"`
	Price       int64  		`json:"price"`
	IsSale      bool   		`json:"is_sale"`
	CreatedAt   string 		`json:"created-at"`
	Image		[]Image 	`json:"images"`
	Rating		float32		`json:"rating"`
	ModifiedAt  string  	`json:"modified-at"`
}