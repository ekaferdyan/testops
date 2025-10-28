package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	// Gunakan 'db' untuk migrasi
	//db.AutoMigrate(&models.User{}, &models.Project{}, &models.TestReport{})

	log.Println("✅ Database connected")

	// 👈 KEMBALIKAN koneksi GORM
	return db
}
