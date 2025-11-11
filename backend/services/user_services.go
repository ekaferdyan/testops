// services/user_service.go
package services

import (
	"errors"
	"time"

	// Semua import yang tidak terpakai akan dihapus
	// "golang.org/x/crypto/bcrypt" - Dihapus
	"sambel-ulek/backend/database"
	"sambel-ulek/backend/models"
	"sambel-ulek/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
	ErrPhoneAlreadyExists = errors.New("nomor handphone sudah terdaftar")
	ErrStatus             = errors.New("status harus active atau inactive")
)

// --- DTO (Data Transfer Object) untuk Payload ---
// Ini tetap harus ada, karena controller Anda menggunakannya
type RegisterRequest struct {
	Email    string `json:"Email"  validate:"required,min=6,max=30,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=2,max=50,name_contains_digits,name_special_character"`
	Phone    string `json:"phone" validate:"required,min=2,max=15,id_phone_not_valid"`
	Status   string `json:"status"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// --- DTO (Data Transfer Object) untuk Response ---
type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"Email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Wrapper Response fungsi nya untuk mengubah model user menjadi response dan bisa dicustomize untuk case ini password tidak kita tampilkan untuk menjaga kerahasiaan
func ToUserResponse(users models.User) UserResponse {
	return UserResponse{
		ID:        users.ID,
		Email:     users.Email,
		Name:      users.Name,
		Phone:     users.Phone,
		Status:    users.Status,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}
}

func RegisterUser(request RegisterRequest) (UserResponse, error) {
	//1. Validasi Bisnis : Check Duplicate Email di Gorm
	var existingUser models.User

	//2. Normalized Phone Number
	request.Phone = utils.NormalizePhone(request.Phone)

	//3. Verify Untuk Email Already Exist
	result := database.ConnectDB().Where("Email = ?", request.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return UserResponse{}, ErrEmailAlreadyExists
	}

	//4. Verify Untuk Phone Already Exist
	result = database.ConnectDB().Where("Phone = ?", request.Phone).First(&existingUser)
	if result.RowsAffected > 0 {
		return UserResponse{}, ErrPhoneAlreadyExists
	}

	//5. Verify Status
	if request.Status != "" {
		if request.Status != "active" && request.Status != "inactive" {
			return UserResponse{}, ErrStatus
		}
	}

	//6. Logika Inti : Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	//7. Logika Inti : Persiapan Data Model
	newUser := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
		Phone:    request.Phone,
		Status:   request.Status,
		// CreatedAt, ID, UpdatedAt akan diisi oleh GORM
	}

	//8. Logika Inti : Simpan ke Database
	if result := database.ConnectDB().Create(&newUser); result.Error != nil {
		return UserResponse{}, result.Error // Error saat menyimpan
	}

	//9. Kembalikan response yang sudah diformat
	return ToUserResponse(newUser), nil
}

func LoginUser(request LoginRequest) (string, error) {

	return "", nil
}
