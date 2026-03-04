package resmodel

import "time"

type DiscountTransactionTarget struct {
	ID            string     `json:"id"` // UUID
	DiscountID    string     `json:"discount_id"`
	MaxTotalQuota int        `json:"max_total_quota"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     string     `json:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *string    `json:"updated_by"`
}

type TransactionDiscountApplied struct {
	ID                          string     `json:"id"` // UUID
	DiscountTransactionTargetID string     `json:"discount_transaction_target_id"`
	TargetID                    string     `json:"target_id"`
	PriceBeforeDiscount         float64    `json:"price_before_discount"`
	TotalDiscount               float64    `json:"total_discount"`
	PriceAfterDiscount          float64    `json:"price_after_discount"`
	IsActive                    bool       `json:"is_active"`
	CreatedAt                   time.Time  `json:"created_at"`
	CreatedBy                   string     `json:"created_by"`
	UpdatedAt                   *time.Time `json:"updated_at"`
	UpdatedBy                   *string    `json:"updated_by"`
}
