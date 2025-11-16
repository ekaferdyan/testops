package dto

import "time"

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
