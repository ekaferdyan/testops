package main

import (
	"log"
	"sambel-ulek/backend/database"
	"sambel-ulek/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	_ = database.ConnectDB()

	routes.SetupAuthRoutes(app)

	// Tambahkan route sederhana untuk testing server hidup
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running! Ready for API requests.")
	})

	// Hidupkan Server di port 3000
	log.Println("Server starting on :3000...")
	log.Fatal(app.Listen(":3000"))
}
