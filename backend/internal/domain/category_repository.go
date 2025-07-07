package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *Category) error
	FindByID(ctx context.Context, id *uuid.UUID) (*Category, error)
	FindAll(ctx context.Context) ([]Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id *uuid.UUID) error
}