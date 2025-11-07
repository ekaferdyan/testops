// routes/auth_routes.go
package routes

import (
	"sambel-ulek/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	api := app.Group("/api") // Grouping semua route di bawah /api

	// Route Public (Tidak Perlu Login)
	api.Post("/register", controllers.RegisterUserHandler)
	api.Post("/login", controllers.Login)

	// Contoh Route yang dilindungi oleh Middleware JWT
	// api.Get("/user", middlewares.JWTAuth(), controllers.GetUser)
}
