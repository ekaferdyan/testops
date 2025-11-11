// controllers/auth_controller.go
package controllers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// Impor "Penerjemah" dan "Pabrik" Validator kita
	platformValidator "sambel-ulek/backend/platform/validator"
	// Impor "Otak Bisnis" kita
	"sambel-ulek/backend/services"
)

// === KODE CONTROLLER ===

func RegisterUserHandler(c *fiber.Ctx) error {
	//1. Parsing
	var request services.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Input JSON tidak valid",
			"errors":  err.Error(),
		})
	}

	//2. Validate Structur untuk mendapatkan message error karena pakai package go
	err := platformValidator.Validate.Struct(request)

	//3. Validasi Input "Fail Fast" (Tugas "Satpam")
	// Menggunakan validator
	if errorMessage := platformValidator.TranslateError(err); errorMessage != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Registrasi gagal",
			"errors":  errorMessage,
		})
	}

	//4. Panggil Service (Jika lolos)
	userResponse, serviceErr := services.RegisterUser(request)

	if serviceErr != nil {
		//Validasi Email Already Exists
		if errors.Is(serviceErr, services.ErrEmailAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Registrasi gagal",
				"errors":  serviceErr.Error(),
			})
		}

		//Validasi Phone Already Exists
		if errors.Is(serviceErr, services.ErrPhoneAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Registrasi gagal",
				"errors":  serviceErr.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan internal",
			"errors":  serviceErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"data":    userResponse,
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
		// if errors.Is(err, services.ErrUserNotFound) { // <-- Mengecek user not found (dummy logic)
		// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ // 404 Not Found
		// 		"status": "error", "message": err.Error(),
		// 	})
		// }
		// if errors.Is(err, services.ErrWrongPassword) { // <-- Mengecek wrong password (dummy logic)
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ // 401 Unauthorized
		// 		"status": "error", "message": err.Error(),
		// 	})
		// }
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
