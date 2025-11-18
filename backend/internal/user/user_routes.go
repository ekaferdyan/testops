package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func NewUserRoutes(db *gorm.DB) {
// 	repo := NewUserRepository(db)
// 	service := NewUserService(repo)
// 	controller := NewUserController(service)
// }

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api") // Grouping semua route di bawah /api

	// ---------------------------------------------------------
	// 1. WIRING / PERAKITAN (Dependency Injection)
	// ---------------------------------------------------------

	// Siapkan Repository (Butuh DB)
	repo := NewUserRepository(db)

	// Siapkan Service (Butuh Repository)
	service := NewUserService(repo)

	// Siapkan Controller (Butuh Service)
	controller := NewUserController(service)

	//endpoint api/v1/register
	api.Post("/v1/register", controller.RegisterUserHandler)

}
