package controller

import (
	"time"

	"github.com/NamDuongkiwi/shopping-backend/database"
	"github.com/NamDuongkiwi/shopping-backend/models"
	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	user.CreatedAt = currentTime

	db := database.Connect()
	defer db.Close()
	_, err = db.Query("INSERT INTO user VALUES (?,?,?,?,?,?,?,?)", user.UserID, user.Name, user.Email, user.Password, user.Avatar, user.Gender, user.Birthday, user.CreatedAt)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	return c.JSON(user)
}	


func GetUserByID(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err!= nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE user_id = ?", id)
	var user models.User
	if rows.Next(){
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.Avatar, &user.Gender, &user.Birthday, &user.CreatedAt)
	}
	return c.JSON(user)
}