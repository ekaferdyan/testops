package mock

import "os/user"

// MockUserRepository adalah struct yang akan menggantikan Repository asli
type MockUser struct {
	IsEmailExistsFunc func(email string) bool
	IsPhoneExistsFunc func(phone string) bool
	CreateUserFunc    func(u *user.User) error
}

// Implementasi interface UserRepository
func (m *MockUser) IsEmailExist(email string) bool {
	return m.IsEmailExistsFunc(email)
}

func (m *MockUser) IsPhoneExists(phone string) bool {
	return m.IsPhoneExistsFunc(phone)
}

func (m *MockUser) CreateUser(u *user.User) error {
	return m.CreateUserFunc(u)
}
