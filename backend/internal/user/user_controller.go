package user

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	// Impor DTO kita
	"sambel-ulek/backend/internal/user/dto"

	// Impor "Penerjemah" dan "Pabrik" Validator kita

	platformValidator "sambel-ulek/backend/platform/validator"
	// Impor "Otak Bisnis" kita
)

// 1. Definisikan Struct Controller (Punya akses ke Service)
type UserController struct {
	service UserService
}

// 2. Constructor Controller (Menerima Service yang sudah ada DB-nya)
func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

// 3. Ubah Handler menjadi METHOD (c *UserController)
// Supaya bisa panggil 'c.service'
func (c *UserController) RegisterUserHandler(ctx *fiber.Ctx) error {
	//1. Parsing
	var request dto.RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Input JSON tidak valid",
			"errors":  err.Error(),
		})
	}

	//2. Validate Structur untuk mendapatkan message error karena pakai package go
	err := platformValidator.Validate.Struct(request)

	//3. Validasi Input "Fail Fast" (Tugas "Satpam")
	// Menggunakan validator
	if errorMessage := platformValidator.TranslateError(err); errorMessage != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Registrasi gagal",
			"errors":  errorMessage,
		})
	}

	//4. Panggil Service (Jika lolos)
	userResponse, serviceErr := c.service.RegisterUser(request)

	if serviceErr != nil {
		//Validasi Email Already Exists
		if errors.Is(serviceErr, ErrEmailAlreadyExists) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Registrasi gagal",
				"errors":  serviceErr.Error(),
			})
		}

		//Validasi Phone Already Exists
		if errors.Is(serviceErr, ErrPhoneAlreadyExists) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Registrasi gagal",
				"errors":  serviceErr.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan internal",
			"errors":  serviceErr.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"data":    userResponse,
	})

}
