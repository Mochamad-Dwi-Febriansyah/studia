package main

import (
	"context"
	"log"
	"studia/backend/internal/domain"
	"studia/backend/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedJurnal(db *gorm.DB) {
	log.Println("memulai menjalankan seed jurnal")
	journalRepo := repository.NewJurnalRepository(db)

	journalToSeed := []domain.Jurnal{
		{
			ID: uuid.New(),
			Activity:    "Mempelajari Clean Architecture",
			Description: "Membuat struktur folder dan dependency injection.",
			Status:      domain.StatusDone, 
		},
		{
			ID:          uuid.New(),
			Activity:    "Implementasi Unit Test",
			Description: "Menulis unit test untuk lapisan usecase menggunakan testify/mock.",
			Status:      domain.StatusDone,
		},
		{
			ID:          uuid.New(),
			Activity:    "Membuat Endpoint API",
			Description: "Menyiapkan handler dan rute menggunakan Gin.",
			Status:      domain.StatusPending,
		},
	}

	for _ , journal := range journalToSeed {
		err := journalRepo.Save(context.Background(), &journal)
		if err != nil {
			log.Printf("gagal menyimpan jurnal %s : %v\n", err.Error(), journal.Activity)
		}else {
			log.Printf("Berhasil menyimpan jurnal : %s\n", journal.Activity)
		}
	}

}