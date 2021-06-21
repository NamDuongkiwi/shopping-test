package controller

import (
	"github.com/NamDuongkiwi/shopping-backend/database"
	"github.com/NamDuongkiwi/shopping-backend/models"
	"github.com/gofiber/fiber/v2"
)


func CreateReview(c *fiber.Ctx)error{
	var review models.Review
	err := c.BodyParser(&review)
	db := database.Connect()
	defer db.Close()
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}

	_, err = db.Query("INSERT INTO review VAlUES (?, ?, ?, ?)", review.ReviewID , review.ProductID, review.Review, review.Rating)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(review)

}

func GetReview(c *fiber.Ctx) error{
	db := database.Connect()
	defer db.Close()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	rows, err := db.Query("SELECT * FROM review WHERE product_id = ?", id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var reviews []models.Review
	for rows.Next(){
		var review models.Review
		err = rows.Scan(&review.ReviewID, &review.ProductID, &review.Review, &review.Rating)
		reviews = append(reviews, review)
	}
	return c.JSON(reviews)
}