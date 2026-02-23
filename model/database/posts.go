package dbmodel

import (
	"time"
)

type Post struct {
	ID        string     `json:"id"`    // UUID
	Title     string     `json:"title"` // UNIQUE
	Content   string     `json:"content"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
}
