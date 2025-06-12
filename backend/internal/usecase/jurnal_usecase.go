package usecase

import (
	"context"
	"studia/backend/internal/domain"
	"time"

	"github.com/google/uuid"
)

type JurnalUsecase interface {
	Create(ctx context.Context, jurnal *domain.Jurnal) error
	FindByID(ctx context.Context, id *uuid.UUID) (*domain.Jurnal, error)
	FindAll(ctx context.Context) ([]domain.Jurnal, error)
	Update(ctx context.Context, jurnal *domain.Jurnal) error
	Delete(ctx context.Context, id *uuid.UUID) error
}

type jurnalUsecase struct {
	jurnalRepo domain.JurnalRepository
	contextTimeout time.Duration
}

func NewJurnalUsecase(jurnalRepo domain.JurnalRepository, timeout time.Duration) JurnalUsecase {
	return &jurnalUsecase{
		jurnalRepo: jurnalRepo,
		contextTimeout: timeout,
	}
}

func (uc *jurnalUsecase) Create(ctx context.Context, jurnal *domain.Jurnal) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	jurnal.Status = domain.StatusPending

	return uc.jurnalRepo.Save(ctx, jurnal)
}

func (uc *jurnalUsecase) FindByID(ctx context.Context, id *uuid.UUID) (*domain.Jurnal, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	
	return uc.jurnalRepo.FindByID(ctx, id)
}

// TAMBAHKAN FUNGSI-FUNGSI BARU DI BAWAH INI

// FindAll menjalankan logika untuk mengambil semua jurnal.
func (uc *jurnalUsecase) FindAll(ctx context.Context) ([]domain.Jurnal, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	return uc.jurnalRepo.FindAll(ctx)
}

// Update menjalankan logika untuk memperbarui jurnal.
func (uc *jurnalUsecase) Update(ctx context.Context, jurnal *domain.Jurnal) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	return uc.jurnalRepo.Update(ctx, jurnal)
}

// Delete menjalankan logika untuk menghapus jurnal.
func (uc *jurnalUsecase) Delete(ctx context.Context, id *uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	return uc.jurnalRepo.Delete(ctx, id)
}