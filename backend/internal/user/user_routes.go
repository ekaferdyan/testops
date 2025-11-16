package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRoutes(db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	controller := NewUserController(service)
}

func SetupAuthRoutes(app *fiber.App) {
	api := app.Group("/api") // Grouping semua route di bawah /api

	//endpoint api/v1/register
	api.Post("/v1/register", registerUserHandler)

}
