package main

import (
	"log"
	"os"
	"studia/backend/pkg/database"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password_mysql_anda"), // Ganti dengan password Anda
		DBName:   getEnv("DB_NAME", "studia_db"),
		Port:     getEnv("DB_PORT", "3306"),
	}
	db := database.NewMySQLDatabase(dbConfig)

	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("Perintah tidak valid. Gunakan: go run ./cmd/seeder [jurnal|user|all]")
		return
	}

	for _, arg := range args {
		switch arg {
		case "jurnal":
			SeedJurnal(db) // Memanggil fungsi dari jurnal_seeder.go
		case "user":
			// Anda bisa membuat user_seeder.go dan memanggil fungsinya di sini
			log.Println("Seeder untuk 'user' belum diimplementasikan.")
		case "all":
			log.Println("Menjalankan semua seeder...")
			SeedJurnal(db)
			// seedUsers(db)
		default:
			log.Printf("Argumen tidak dikenal: %s", arg)
		}
	}
}
