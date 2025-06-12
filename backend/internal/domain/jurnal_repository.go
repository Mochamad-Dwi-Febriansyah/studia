package domain

import (
	"context"

	"github.com/google/uuid"
)

type JurnalRepository interface {
	Save(ctx context.Context, jurnal *Jurnal) error

	FindByID(ctx context.Context, id *uuid.UUID) (*Jurnal, error)

	FindAll(ctx context.Context) ([]Jurnal, error)

	Update(ctx context.Context, jurnal *Jurnal) error

	Delete(ctx context.Context, id *uuid.UUID) error
}