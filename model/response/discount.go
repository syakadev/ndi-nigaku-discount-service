package resmodel

import "time"

type Discount struct {
	ID        string     `json:"id"` // UUID
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Value     float64    `json:"value"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	Target    string     `json:"target"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}
