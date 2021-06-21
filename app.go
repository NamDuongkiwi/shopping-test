package main

import (
	"github.com/NamDuongkiwi/shopping-backend/routes"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	//fmt.Println(controller.GetRating(1))
	productRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&productRouter)

	reviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&reviewRouter)
	
	userRouter := app.Group("/api/user")
	routes.ConfigUserRouter(&userRouter)

	app.Listen(":3000")
}