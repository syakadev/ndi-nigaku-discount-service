package resmodel

import "time"

type DiscountProductTarget struct {
	ID            string     `json:"id"` // UUID
	DiscountID    string     `json:"discount_id"`
	TargetID      string     `json:"target_id"`
	MaxTotalQuota int        `json:"max_total_quota"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     string     `json:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *string    `json:"updated_by"`
}

type ProductDiscountApplied struct {
	ID string `json:"id"` // UUID
	ProductDiscountID string `json:"product_discount_id"`
	CustomerID	string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     string     `json:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *string    `json:"updated_by"`
}