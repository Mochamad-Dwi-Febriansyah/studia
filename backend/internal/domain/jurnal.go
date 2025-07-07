package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusLevel string

const (
	StatusPending StatusLevel = "pending"
	StatusRejected StatusLevel = "rejected"
	StatusDone StatusLevel = "done"
) 

type Jurnal struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Slug string `gorm:"type:varchar(191);null;uniqueIndex" json:"slug"`
	Activity string `gorm:"not null" json:"activity"`
	Description string  `gorm:"type:text;not null" json:"description"` 
	Status StatusLevel `gorm:"type:varchar(15)" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"` 
}