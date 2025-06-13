package repository

import (
	"context"
	"studia/backend/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type jurnalRepository struct {
	db *gorm.DB
}

func NewJurnalRepository(db *gorm.DB) domain.JurnalRepository {
	return &jurnalRepository{
		db : db,
	}
}

func (r *jurnalRepository) Save(ctx context.Context, jurnal *domain.Jurnal) error {
	return r.db.WithContext(ctx).Create(jurnal).Error
}

func (r *jurnalRepository) FindByID(ctx context.Context, id *uuid.UUID) (*domain.Jurnal, error) {
	var jurnal domain.Jurnal
	err := r.db.WithContext(ctx).First(&jurnal, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}
 
func (r *jurnalRepository) FindAll(ctx context.Context) ([]domain.Jurnal, error) {
	var jurnals []domain.Jurnal 
	err := r.db.WithContext(ctx).Find(&jurnals).Error
	if err != nil {
		return nil, err
	}
	return jurnals, nil
}
 
func (r *jurnalRepository) Update(ctx context.Context, jurnal *domain.Jurnal) error { 
	return r.db.WithContext(ctx).Save(jurnal).Error
}
 
func (r *jurnalRepository) Delete(ctx context.Context, id *uuid.UUID) error { 
	return r.db.WithContext(ctx).Delete(&domain.Jurnal{}, id).Error
}
