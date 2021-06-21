package controller

import (
	//"fmt"
	"fmt"
	"time"

	"github.com/NamDuongkiwi/shopping-backend/database"
	"github.com/NamDuongkiwi/shopping-backend/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
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
		product.Image = GetImageProduct(product.ID)
		product.Rating = 4.5
		products = append(products, product)
	}
	return c.JSON(products)
}

func GetImageProduct(id int) []string {
	db := database.Connect()
	defer db.Close()

	rows, erro := db.Query("SELECT url FROM product_image WHERE product_id = ?", id)
	if erro != nil {
		panic(erro.Error())
	}
	var images []string
	for rows.Next(){
		var image string
		if err := rows.Scan(&image); err != nil {
			panic(err.Error())
		}
		images = append(images, image)
	}
	return images
}

func CreateProduct(c *fiber.Ctx) error{
	
	db := database.Connect()
	defer db.Close()
	var product models.Product
	err := c.BodyParser(&product)
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	_ , err = db.Query("INSERT INTO product VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Category_id, product.Price, currentTime, currentTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	_, err = db.Query("INSERT INTO price_update VALUES(?, ?, ?)", product.ID, currentTime, product.Price)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"product" : product,
	})

}

func GetProductbyID(c *fiber.Ctx) error{
	db:=database.Connect()
	defer db.Close()
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	row, err := db.Query("SELECT * FROM product WHERE id = ?", Id)
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "id not found",
		})
	}
	var product models.Product
	if row.Next(){
		row.Scan(&product.ID, &product.Name, &product.Category_id, &product.Price, &product.IsSale, &product.CreatedAt, &product.ModifiedAt)
	}
	product.Image = GetImageProduct(product.ID)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx)error{
	db:=database.Connect()
	defer db.Close()
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	_, err = db.Query("DELETE FROM product_image WHERE product_id = ?", Id)
	if err != nil{
		fmt.Println(err.Error())
	}

	_, err = db.Query("DELETE FROM price_update WHERE product_id = ?", Id)
	if err != nil{
		fmt.Println(err.Error())
	}
	_, err = db.Query("DELETE FROM product WHERE id = ?", Id)
	if err != nil{
		fmt.Println(err.Error())
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"product" : Id,
	})
}