package repository

import (
	"context"
	"studia/backend/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Save(ctx context.Context, category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindByID(ctx context.Context, id *uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).First(category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Category{}, id).Error
}	