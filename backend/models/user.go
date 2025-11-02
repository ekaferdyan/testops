package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}

// Tambahkan method ini jika nama skema Anda BUKAN "public"
// Ganti "data" dengan nama skema Anda yang sebenarnya.
func (User) TableName() string {
	return "DATA.users" // <--- GANTI 'data' DENGAN NAMA SKEMA ANDA
}
