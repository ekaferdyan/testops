// services/user_service.go
package services

import (
	"errors"
	"time"
	// Semua import yang tidak terpakai akan dihapus
	// "golang.org/x/crypto/bcrypt" - Dihapus
	// "testops-dashboard/backend/config" - Dihapus
	// "testops-dashboard/backend/models" - Dihapus
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
	ErrUserNotFound       = errors.New("pengguna tidak ditemukan")
	ErrWrongPassword      = errors.New("kata sandi salah")
)

// --- DTO (Data Transfer Object) untuk Request ---
// Ini tetap harus ada, karena controller Anda menggunakannya
type RegisterRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// --- DTO (Data Transfer Object) untuk Response ---
type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Logika Bisnis (SIMPLIFIED / DUMMY) ---

func RegisterUser(request RegisterRequest) (UserResponse, error) {

	// --- SIMULASI LOGIKA ---

	// Jika email-nya "tes@error.com", kita simulasikan email sudah ada.
	if request.Email == "error@fail.com" {
		return UserResponse{}, ErrEmailAlreadyExists
	}

	// Jika tidak, anggap sukses.
	dummyResponse := UserResponse{
		ID:        uint(time.Now().Unix()), // ID palsu
		Email:     request.Email,
		CreatedAt: time.Now(),
	}

	// Sukses: Kembalikan data palsu dan error nil
	return dummyResponse, nil
}

func LoginUser(request LoginRequest) (string, error) {

	// --- SIMULASI LOGIKA ---

	// 1. Simulasikan Not Found
	if request.Email == "notfound@user.com" {
		return "", ErrUserNotFound
	}

	// 2. Simulasikan Wrong Password
	if request.Password == "salah" {
		return "", ErrWrongPassword
	}

	// 3. Simulasikan Sukses (Kembalikan string token palsu)
	tokenPalsu := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.SflKxwR"

	// Sukses: Kembalikan token palsu dan error nil
	return tokenPalsu, nil
}
