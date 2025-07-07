package usecase

import (
	"context"
	"studia/backend/internal/domain"
	"time"

	"github.com/google/uuid"
)

type CategoryUsecase interface {
	Create(ctx context.Context, category *domain.Category) error
	FindByID(ctx context.Context, id *uuid.UUID) (*domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id *uuid.UUID) error
}

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
	contextTimeout            time.Duration
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, timeout time.Duration) CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
		contextTimeout:            timeout,
	}
}

func (uc *categoryUsecase) Create(ctx context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	category.StatusView = domain.StatusDraft

	return uc.categoryRepo.Save(ctx, category)
}

func (uc *categoryUsecase) FindAll(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	categories, err := uc.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (uc *categoryUsecase) FindByID(ctx context.Context, id *uuid.UUID) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	category, err := uc.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *categoryUsecase) Update(ctx context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if category.StatusView == "" {
		category.StatusView = domain.StatusDraft
	}

	return uc.categoryRepo.Update(ctx, category)
}

func (uc *categoryUsecase) Delete(ctx context.Context, id *uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	return uc.categoryRepo.Delete(ctx, id)
}