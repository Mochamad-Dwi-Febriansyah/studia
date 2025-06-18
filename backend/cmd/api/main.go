package main

import (
	"log"
	"os"
	"studia/backend/internal/handler/http"
	"studia/backend/internal/repository"
	"studia/backend/internal/usecase"
	"studia/backend/pkg/database" // Menggunakan paket database yang sudah kita buat
	"time"

	"github.com/gin-gonic/gin"
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
		Password: getEnv("DB_PASSWORD", "123"),  
		DBName:   getEnv("DB_NAME", "db_studia"),
		Port:     getEnv("DB_PORT", "3306"),
	} 

	db := database.NewMySQLDatabase(dbConfig)
 
	timeoutContext := 2 * time.Second

	jurnalRepo := repository.NewJurnalRepository(db)
	jurnalUsecase := usecase.NewJurnalUsecase(jurnalRepo, timeoutContext)

	r := gin.Default() 

	api := r.Group("api")
	
	http.NewJurnalHandler(api, jurnalUsecase)
 
	log.Println("Server akan berjalan di port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
