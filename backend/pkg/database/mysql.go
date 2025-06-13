package database

import (
	"fmt"
	"log"
	"studia/backend/internal/domain"

	"gorm.io/driver/mysql" // Ganti driver ke mysql
	"gorm.io/gorm"
)

// Config menampung semua konfigurasi untuk koneksi database.
type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

// NewMySQLDatabase membuat koneksi baru ke database MySQL dan menjalankan AutoMigrate.
func NewMySQLDatabase(config Config) *gorm.DB {
	// DSN (Data Source Name) untuk MySQL memiliki format yang berbeda.
	// user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	// Ganti driver dari postgres.Open ke mysql.Open
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Menjalankan migrasi otomatis, ini akan tetap bekerja
	db.AutoMigrate(&domain.Jurnal{})
	log.Println("Koneksi database MySQL dan migrasi berhasil.")

	return db
}
