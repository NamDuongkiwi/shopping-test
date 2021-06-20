package controller

import (
	"database/sql"
	"github.com/NamDuongkiwi/shopping-backend/models"
	"github.com/gofiber/fiber/v2"
)

type dbStore struct {
	db *sql.DB
}


func (store *dbStore) GetAllProducts(c *fiber.Ctx) error {
	rows, err := store.db.Query("SELECT * FROM product")
	if err != nil {
		return err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next(){
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category_id, &product.Price, &product.IsSale, &product.CreatedAt, &product.Rating)
		if err != nil {
            return err
        }
		products = append(products, product)
	}
	return c.JSON(products)
}
