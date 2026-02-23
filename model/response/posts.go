package resmodel

import (
	"time"
)

type Post struct {
	ID        string     `json:"id"` // UUID
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}
