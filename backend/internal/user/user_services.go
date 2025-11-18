package user

import (
	"errors"
	"sambel-ulek/backend/internal/user/dto"
	"sambel-ulek/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
	ErrPhoneAlreadyExists = errors.New("nomor handphone sudah terdaftar")
	ErrStatus             = errors.New("status harus active atau inactive")
)

// Service Struct
type UserService interface {
	RegisterUser(request dto.RegisterRequest) (dto.UserResponse, error)
}

// IMPLEMENTATION
type userService struct {
	repo UserRepository
}

// Constructor DI
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// Wrapper Response fungsi nya untuk mengubah model user menjadi response dan bisa dicustomize untuk case ini password tidak kita tampilkan untuk menjaga kerahasiaan
func ToUserResponse(users User) dto.UserResponse {
	return dto.UserResponse{
		ID:        users.ID,
		Email:     users.Email,
		Name:      users.Name,
		Phone:     users.Phone,
		Status:    users.Status,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}
}

func (s *userService) RegisterUser(request dto.RegisterRequest) (dto.UserResponse, error) {
	//2. Normalized Phone Number
	request.Phone = utils.NormalizePhone(request.Phone)

	//3. Verify Untuk Email Already Exist
	if s.repo.IsEmailExists(request.Email) {
		return dto.UserResponse{}, ErrEmailAlreadyExists
	}

	//4. Verify Untuk Phone Already Exist
	if s.repo.IsPhoneExists(request.Phone) {
		return dto.UserResponse{}, ErrPhoneAlreadyExists
	}

	//5. Verify Status
	if request.Status != "" {
		if request.Status != "active" && request.Status != "inactive" {
			return dto.UserResponse{}, ErrStatus
		}
	}

	//6. Logika Inti : Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	//7. Logika Inti : Persiapan Data Model
	newUser := &User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
		Phone:    request.Phone,
		Status:   request.Status,
		// CreatedAt, ID, UpdatedAt akan diisi oleh GORM
	}

	//8. Logika Inti : Simpan ke Database
	if err := s.repo.CreateUser(newUser); err != nil {
		return dto.UserResponse{}, err
	}

	//9. Kembalikan response yang sudah diformat
	return ToUserResponse(*newUser), nil
}
