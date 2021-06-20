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
	productRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&productRouter)
	app.Listen(":3000")
}