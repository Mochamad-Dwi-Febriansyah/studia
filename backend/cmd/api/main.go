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
		Host:     getEnv("DB_HOST", "localhost"), // <-- Diubah untuk membaca env var
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "Terserah123."), // <-- Ganti 'secret' dengan password default jika perlu
		DBName:   getEnv("DB_NAME", "db_studia"),
		Port:     getEnv("DB_PORT", "3306"),
	} 
	db := database.NewMySQLDatabase(dbConfig)
 
	timeoutContext := 2 * time.Second

	jurnalRepo := repository.NewJurnalRepository(db)
	jurnalUsecase := usecase.NewJurnalUsecase(jurnalRepo, timeoutContext)
	jurnalHandler := http.NewJurnalHandler(jurnalUsecase)
 
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		jurnalRoutes := v1.Group("/jurnals")
		{
			jurnalRoutes.POST("/", jurnalHandler.Create)
			jurnalRoutes.GET("/", jurnalHandler.FindAll)
			jurnalRoutes.GET("/:id", jurnalHandler.FindByID)
			jurnalRoutes.PUT("/:id", jurnalHandler.Update)
			jurnalRoutes.DELETE("/:id", jurnalHandler.Delete)
		}
	}
 
	log.Println("Server akan berjalan di port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
