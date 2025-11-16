package user

import (
	"errors"
	"sambel-ulek/backend/internal/user/dto"
	"sambel-ulek/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// --- Definisi Error Bisnis (Tetap ada, untuk uji coba error handling) ---
var (
	errEmailAlreadyExists = errors.New("email sudah terdaftar")
	errPhoneAlreadyExists = errors.New("nomor handphone sudah terdaftar")
	errStatus             = errors.New("status harus active atau inactive")
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
func ToUserResponse(users user) dto.UserResponse {
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
		return dto.UserResponse{}, errEmailAlreadyExists
	}

	//4. Verify Untuk Phone Already Exist
	if s.repo.IsPhoneExists(request.Phone) {
		return dto.UserResponse{}, errPhoneAlreadyExists
	}

	//5. Verify Status
	if request.Status != "" {
		if request.Status != "active" && request.Status != "inactive" {
			return dto.UserResponse{}, errStatus
		}
	}

	//6. Logika Inti : Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	//7. Logika Inti : Persiapan Data Model
	newUser := user{
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
	return ToUserResponse(newUser), nil
}
