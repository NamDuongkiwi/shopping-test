package controller

import (
	"github.com/NamDuongkiwi/shopping-backend/models"
	"github.com/NamDuongkiwi/shopping-backend/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
)




func GetAllCategory(c *fiber.Ctx) error{
	db := database.Connect()
	defer db.Close()
	rows, erro := db.Query("SELECT * FROM category")
	if erro != nil {
		panic(erro.Error())
	}
	var categorys []models.Category
	for rows.Next(){
		var category models.Category
		err:= rows.Scan(&category.CategoryId, &category.Name)
		if err != nil {
            panic("ko co gi ca")
        }
		categorys = append(categorys, category)
	}
	return c.JSON(categorys)
}
func GetAllProducts(c *fiber.Ctx) error{
	db:=database.Connect()
	defer db.Close()
	rows, erro := db.Query("SELECT * FROM product")
	if erro != nil {
		panic(erro.Error())
	}
	var products []models.Product
	for rows.Next(){
		var product models.Product
		err:= rows.Scan(&product.ID, &product.Name, &product.Category_id, &product.Price, &product.IsSale, &product.CreatedAt, &product.ModifiedAt)
		if err != nil {
            panic(err.Error())
        }
		product.Rating = 4.5
		products = append(products, product)
	}
	return c.JSON(products)
}