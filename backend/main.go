package main

import (
	"log"
	"sambel-ulek/backend/database"
	"sambel-ulek/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// init() akan berjalan otomatis SEBELUM main()
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[WARN]: Error memuat file .env.")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// 1. Panggil ConnectDB() SATU KALI dan simpan koneksinya
	database.ConnectDB()

	routes.SetupAuthRoutes(app)

	// Tambahkan route sederhana untuk testing server hidup
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running! Ready for API requests.")
	})

	// Hidupkan Server di port 3000
	log.Println("Server starting on :3000...")
	log.Fatal(app.Listen(":3000"))
}
