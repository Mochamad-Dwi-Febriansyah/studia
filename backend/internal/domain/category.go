package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusView string

const (
	StatusDraft   StatusView = "draft"
	StatusPublished StatusView = "published" 
)

type Category struct {
	ID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Slug string `gorm:"null;uniqueIndex" json:"slug"`  
	Description string `gorm:"type:text" json:"description"`
	StatusView StatusView `gorm:"type:varchar(15)" json:"status_view"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}