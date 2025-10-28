// controllers/auth_controller.go
package controllers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// Impor "Penerjemah" dan "Pabrik" Validator kita
	platformValidator "testops-dashboard/backend/platform/validator"
	// Impor "Otak Bisnis" kita
	"testops-dashboard/backend/services"
)

// === KODE CONTROLLER ===

func Register(c *fiber.Ctx) error {
	var request services.RegisterRequest

	// 2. Parsing
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Input JSON tidak valid", "data": err.Error(),
		})
	}

	// 3. Validasi Format
	if err := platformValidator.Validate.Struct(request); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			translatedErrors := platformValidator.TranslateError(validationErrors)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error", "message": "Validasi gagal", "data": translatedErrors,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Kesalahan validasi", "data": err.Error(),
		})
	}

	// 4. Delegasikan ke "Otak Bisnis" (Service Layer) - Menggunakan logic dummy/simpel
	userResponse, err := services.RegisterUser(request)
	if err != nil {
		// Cek apakah ini error BISNIS yang kita kenal
		if errors.Is(err, services.ErrEmailAlreadyExists) { // <-- Mengecek error dari service dummy
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // 400 Bad Request
				"status": "error", "message": err.Error(), // Pesan: "email sudah terdaftar"
			})
		}
		// Jika bukan, ini error TEKNIS
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ // 500 Internal Error
			"status": "error", "message": "Server gagal mendaftarkan pengguna",
		})
	}

	// 5. Sukses!
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{ // 201 Created
		"status":  "success",
		"message": "User registered successfully",
		"data":    userResponse, // Data dummy akan dikembalikan
	})
}

func Login(c *fiber.Ctx) error {
	var request services.LoginRequest

	// 2. Parsing
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Input JSON tidak valid", "data": err.Error(),
		})
	}

	// 3. Validasi Format
	if err := platformValidator.Validate.Struct(request); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			translatedErrors := platformValidator.TranslateError(validationErrors)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error", "message": "Validasi gagal", "data": translatedErrors,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Kesalahan validasi", "data": err.Error(),
		})
	}

	// 4. Delegasikan ke "Otak Bisnis" (Service Layer)
	token, err := services.LoginUser(request)
	if err != nil {
		// Cek error BISNIS
		if errors.Is(err, services.ErrUserNotFound) { // <-- Mengecek user not found (dummy logic)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ // 404 Not Found
				"status": "error", "message": err.Error(),
			})
		}
		if errors.Is(err, services.ErrWrongPassword) { // <-- Mengecek wrong password (dummy logic)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ // 401 Unauthorized
				"status": "error", "message": err.Error(),
			})
		}
		// Error TEKNIS (dummy logic)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ // 500 Internal Error
			"status": "error", "message": "Server gagal memproses login",
		})
	}

	// 5. Sukses!
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ // 200 OK
		"status":  "success",
		"message": "Login successful",
		"token":   token, // Token palsu akan dikembalikan
	})
}
