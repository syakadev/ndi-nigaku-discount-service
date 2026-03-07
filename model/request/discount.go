package reqmodel

import "time"

type CreateDiscount struct {
	Name       string    `json:"name" validate:"required"`
	Type       string    `json:"type" validate:"required"`
	Value      float64   `json:"value" validate:"required"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
	Target     string    `json:"target" validate:"required"`
	AuthUserID string    `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateDiscount struct {
	ID         string    `json:"id" validate:"required" swaggerignore:"true"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Value      float64   `json:"value"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Target     string    `json:"target"`
	IsActive   bool      `json:"is_active"`
	AuthUserID string    `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}
