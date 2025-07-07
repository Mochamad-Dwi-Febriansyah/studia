package database

import (
	"fmt"
	"log"
	"studia/backend/internal/domain"

	"gorm.io/driver/mysql" // Ganti driver ke mysql
	"gorm.io/gorm"
)
 
type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}
 
func NewMySQLDatabase(config Config) *gorm.DB { 
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
 
	db.AutoMigrate(
		&domain.Jurnal{},
		&domain.Category{},
	)

	log.Println("Koneksi database MySQL dan migrasi berhasil.")

	return db
}
