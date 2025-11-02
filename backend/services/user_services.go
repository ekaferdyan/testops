// services/user_service.go
package services

import (
	"errors"
	"time"

	// Semua import yang tidak terpakai akan dihapus
	// "golang.org/x/crypto/bcrypt" - Dihapus
	"testops-dashboard/backend/database"
	"testops-dashboard/backend/models"
	"testops-dashboard/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	ErrUsernameAlreadyExists = errors.New("username sudah terdaftar")
	ErrUsernameSpecialChar   = errors.New("username tidak boleh mengandung spesial karakter")
)

// --- DTO (Data Transfer Object) untuk Payload ---
// Ini tetap harus ada, karena controller Anda menggunakannya
type RegisterRequest struct {
	Username string `json:"username"  validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginRequest struct {
	Username string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// --- DTO (Data Transfer Object) untuk Response ---
type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

// Wrapper Response fungsi nya untuk mengubah model user menjadi response dan bisa dicustomize untuk case ini password tidak kita tampilkan untuk menjaga kerahasiaan
func ToUserResponse(users models.User) UserResponse {
	return UserResponse{
		ID:        users.ID,
		Username:  users.Username,
		CreatedAt: users.CreatedAt,
	}
}

func RegisterUser(request RegisterRequest) (UserResponse, error) {
	//1. Validasi Bisnis : Check Duplicate Username di Gorm
	var existingUser models.User

	//1. Verify Untuk Username Already Exist
	result := database.ConnectDB().Where("username = ?", request.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		return UserResponse{}, ErrUsernameAlreadyExists
	}

	//2. Verify Username Mengandung Special Karakter
	// Log hasil dari fungsi utilitas
	if utils.ContainsSpecialCharacters(request.Username) {
		return UserResponse{}, ErrUsernameSpecialChar
	}

	//3. Logika Inti : Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	//4. Logika Inti : Persiapan Data Model
	newUser := models.User{
		Username: request.Username,
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
