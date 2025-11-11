package user

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	// Impor "Penerjemah" dan "Pabrik" Validator kita
	platformValidator "sambel-ulek/backend/platform/validator"
	// Impor "Otak Bisnis" kita
)

// === KODE CONTROLLER ===

func RegisterUserHandler(c *fiber.Ctx) error {
	//1. Parsing
	var request registerRequest
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
	userResponse, serviceErr := RegisterUser(request)

	if serviceErr != nil {
		//Validasi Email Already Exists
		if errors.Is(serviceErr, errEmailAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Registrasi gagal",
				"errors":  serviceErr.Error(),
			})
		}

		//Validasi Phone Already Exists
		if errors.Is(serviceErr, errPhoneAlreadyExists) {
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
