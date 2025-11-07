package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(50);unique;not null"`
	Name      string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(15);unique;not null"`
	Status    string `gorm:"type:varchar(10);default:inactive;not null"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Tambahkan method ini jika nama skema Anda BUKAN "public"
// Ganti "data" dengan nama skema Anda yang sebenarnya.
func (User) TableName() string {
	return "AUTHENTICATION.users" // <--- GANTI 'data' DENGAN NAMA SKEMA ANDA
}
