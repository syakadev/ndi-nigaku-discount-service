package resmodel

import "time"

type DiscountProductTarget struct {
	ID                  string     `json:"id"` // UUID
	DiscountID          string     `json:"discount_id"`
	TargetID            string     `json:"target_id"`
	ProductName         string     `json:"product_name"`
	MaxTotalQuota       int        `json:"max_total_quota"`
	PriceBeforeDiscount float64    `json:"price_before_discount"`
	TotalDiscount       float64    `json:"total_discount"`
	PriceAfterDiscount  float64    `json:"price_after_discount"`
	IsActive            bool       `json:"is_active"`
	CreatedAt           time.Time  `json:"created_at"`
	CreatedBy           string     `json:"created_by"`
	UpdatedAt           *time.Time `json:"updated_at"`
	UpdatedBy           *string    `json:"updated_by"`
}

type ProductDiscountApplied struct {
	ID                string     `json:"id"` // UUID
	ProductDiscountID string     `json:"product_discount_id"`
	CustomerID        string     `json:"customer_id"`
	CustomerName      string     `json:"customer_name"`
	TransactionDate   time.Time  `json:"transaction_date"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
}
