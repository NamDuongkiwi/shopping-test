package controller

import (
	"time"

	"github.com/NamDuongkiwi/shopping-backend/database"
	"github.com/NamDuongkiwi/shopping-backend/models"
	//	"github.com/NamDuongkiwi/shopping-backend/models"
	"github.com/gofiber/fiber/v2"
)


func ChangePrice(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	var price models.PriceUpdade
	err = c.BodyParser(&price)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	price.ProductID = id
	price.UpdatedAt = currentTime
	db := database.Connect()
	defer db.Close()
	_, err = db.Query("INSERT INTO price_update VALUES (?, ?, ?)", price.ProductID, price.UpdatedAt, price.Price)
	_, err = db.Query("UPDATE product SET price = ? WHERE id = ?", price.Price, id)
	
	return c.JSON(price)
}

func GetPriceHistory(c *fiber.Ctx) error{
	db := database.Connect()
	defer db.Close()
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	rows, err := db.Query("SELECT * FROM price_update WHERE product_id = ?", id)

	var priceUpdates []models.PriceUpdade

	for rows.Next(){
		var priceUpdate models.PriceUpdade
		err = rows.Scan(&priceUpdate.ProductID, &priceUpdate.UpdatedAt, &priceUpdate.Price)
		priceUpdates = append(priceUpdates, priceUpdate)
	}
	return c.JSON(priceUpdates)
}