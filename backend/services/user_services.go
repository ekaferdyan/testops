// services/user_service.go
package services

import (
	"errors"
	"time"

	// Semua import yang tidak terpakai akan dihapus
	// "golang.org/x/crypto/bcrypt" - Dihapus
	"sambel-ulek/backend/database"
	"sambel-ulek/backend/models"

	"golang.org/x/crypto/bcrypt"
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
	ErrEmailNotValid      = errors.New("email tidak valid")
	ErrPhoneAlreadyExists = errors.New("nomor handphone sudah terdaftar")
	ErrNamespecialChar    = errors.New("nama tidak boleh mengandung spesial karakter")
	ErrStatus             = errors.New("status harus active atau inactive")
	ErrNameAngka          = errors.New("nama tidak boleh mengandung angka")
	ErrPhoneNotValid      = errors.New("nomor handphone tidak valid")
	ErrPhoneSpecialChar   = errors.New("nomor handphone tidak boleh mengandung spesial karakter")
	ErrPhoneHuruf         = errors.New("nomor handphone tidak boleh mengandung huruf")
)

// --- DTO (Data Transfer Object) untuk Payload ---
// Ini tetap harus ada, karena controller Anda menggunakannya
type RegisterRequest struct {
	Email    string `json:"Email"  validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Phone    string `json:"phone" validate:"required,min=2,max=15"`
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
	UpdateAt  time.Time `json:"update_at"`
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
		UpdateAt:  users.UpdateAt,
	}
}

func RegisterUser(request RegisterRequest) (UserResponse, error) {
	//1. Validasi Bisnis : Check Duplicate Email di Gorm
	var existingUser models.User

	//1. Verify Untuk Email Already Exist
	result := database.ConnectDB().Where("Email = ?", request.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return UserResponse{}, ErrEmailAlreadyExists
	}

	//2. Verify Email Mengandung Special Karakter
	// Log hasil dari fungsi utilitas
	// if utils.ContainsSpecialCharacters(request.Email) {
	// 	return UserResponse{}, ErrEmailSpecialChar
	// }

	//3. Logika Inti : Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	//4. Logika Inti : Persiapan Data Model
	newUser := models.User{
		Email:    request.Email,
		Password: string(hashedPassword),
		// CreatedAt, ID akan diisi oleh GORM
	}

	// 5. Logika Inti : Simpan ke Database
	if result := database.ConnectDB().Create(&newUser); result.Error != nil {
		return UserResponse{}, result.Error // Error saat menyimpan
	}

	// 6. Kembalikan response yang sudah diformat
	return ToUserResponse(newUser), nil
}

func LoginUser(request LoginRequest) (string, error) {

	return "", nil
}
