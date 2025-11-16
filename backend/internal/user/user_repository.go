package user

import (
	"gorm.io/gorm"
)

// INTERFACE REPOSITORY
type UserRepository interface {
	IsEmailExists(email string) bool
	IsPhoneExists(phone string) bool
	CreateUser(u user) error
}

// -----------------------------------------------------
//  2. STRUCT yang implement interface
//     (Dependency Injection: db masuk ke repository)
//
// -----------------------------------------------------
type userRepository struct {
	db *gorm.DB
}

// Constructor (ini dipanggil dari userRoutes)
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) IsEmailExists(email string) bool {
	//Deskripsi Variable
	var count int64
	r.db.Model(&user{}).Where("phone = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) IsPhoneExists(phone string) bool {
	//Deskripsi Variable
	var count int64
	r.db.Model(&user{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func (r *userRepository) CreateUser(u user) error {
	result := r.db.Create(&u)
	return result.Error
}
