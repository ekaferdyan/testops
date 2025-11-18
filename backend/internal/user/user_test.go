// sambel-ulek/backend/internal/user/service_test.go
package user

import (
	"errors"
	"sambel-ulek/backend/internal/user/dto"
	"sambel-ulek/backend/internal/user/mock" // Import mock repository kita
	"testing"
)

// Catatan: Asumsikan entity/model user bernama 'User' dan dipanggil dari 'user.User'

func TestRegisterUser(t *testing.T) {

	// 1. Buat data Request dummy yang valid
	validReq := dto.RegisterRequest{
		Name:     "Joko Santoso",
		Email:    "joko@test.com",
		Phone:    "081234567890",
		Password: "password123",
		Status:   "active",
	}

	// 2. Definisi Test Cases
	tests := []struct {
		name string
		req  dto.RegisterRequest
		// Config Mock Repository
		mockEmailExist func(email string) bool
		mockPhoneExist func(phone string) bool
		mockCreate     func(u *User) error // Sesuaikan tipe data
		// Ekspektasi
		expectedError error
	}{
		{
			name:           "Success - New User Registration (Active)",
			req:            validReq,
			mockEmailExist: func(email string) bool { return false },
			mockPhoneExist: func(phone string) bool { return false },
			mockCreate:     func(u *User) error { return nil },
			expectedError:  nil,
		},
		{
			name:           "Failure - Email Already Exists",
			req:            validReq,
			mockEmailExist: func(email string) bool { return true }, // <-- Email sudah ada
			mockPhoneExist: func(phone string) bool { return false },
			mockCreate:     func(u *User) error { return nil },
			expectedError:  ErrEmailAlreadyExists, // <-- Error yang diharapkan
		},
		{
			name:           "Failure - Phone Already Exists",
			req:            validReq,
			mockEmailExist: func(email string) bool { return false },
			mockPhoneExist: func(phone string) bool { return true }, // <-- Phone sudah ada
			mockCreate:     func(u *User) error { return nil },
			expectedError:  ErrPhoneAlreadyExists,
		},
		{
			name:           "Failure - Invalid Status Value",
			req:            dto.RegisterRequest{Name: "Test", Email: "a@b.com", Phone: "0", Password: "p", Status: "pending"}, // Status yang tidak valid
			mockEmailExist: func(email string) bool { return false },
			mockPhoneExist: func(phone string) bool { return false },
			mockCreate:     func(u *User) error { return nil },
			expectedError:  ErrStatus, // <-- Error yang diharapkan
		},
		{
			name:           "Failure - Database Creation Error",
			req:            validReq,
			mockEmailExist: func(email string) bool { return false },
			mockPhoneExist: func(phone string) bool { return false },
			mockCreate:     func(u *User) error { return errors.New("simulasi error DB") }, // <-- Simulasi error DB
			expectedError:  errors.New("simulasi error DB"),
		},
	}

	// 3. Loop melalui test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Inisialisasi Mock Repository dengan fungsi yang kita definisikan di test case
			mockRepo := &mock.MockUserRepository{
				IsEmailExistsFunc: tt.mockEmailExist,
				IsPhoneExistsFunc: tt.mockPhoneExist,
				CreateUserFunc:    tt.mockCreate,
			}

			// Inisialisasi Service dengan Mock Repository (Dependency Injection)
			service := NewUserService(mockRepo)

			// Jalankan fungsi yang diuji
			_, actualErr := service.RegisterUser(tt.req)

			// Cek hasil
			if tt.expectedError != nil {
				// Jika error diharapkan
				if !errors.Is(actualErr, tt.expectedError) && actualErr.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedError, actualErr)
				}
			} else {
				// Jika sukses diharapkan
				if actualErr != nil {
					t.Errorf("Expected nil error, but got %v", actualErr)
				}
			}
		})
	}
}
