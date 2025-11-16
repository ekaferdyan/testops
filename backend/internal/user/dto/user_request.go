package dto

// --- DTO (Data Transfer Object) untuk Payload ---
// Ini tetap harus ada, karena controller Anda menggunakannya
type RegisterRequest struct {
	Email    string `json:"Email"  validate:"required,min=6,max=30,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=2,max=50,name_contains_digits,name_special_character"`
	Phone    string `json:"phone" validate:"required,min=2,max=15,id_phone_not_valid"`
	Status   string `json:"status"`
}
