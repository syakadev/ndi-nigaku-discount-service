package reqmodel

type CreateDiscountTransactionTarget struct {
	DiscountID    string `json:"discount_id" validate:"required"`
	MaxTotalQuota int    `json:"max_total_quota" validate:"required"`
	IsActive      bool   `json:"is_active" validate:"required"`
	AuthUserID    string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateDiscountTransactionTarget struct {
	ID            string `json:"id" validate:"required" swaggerignore:"true"`
	DiscountID    string `json:"discount_id"`
	MaxTotalQuota int    `json:"max_total_quota"`
	IsActive      bool   `json:"is_active"`
	AuthUserID    string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type CreateTransactionDiscountApplied struct {
	DiscountTransactionTargetID string  `json:"discount_transaction_target_id" validate:"required"`
	TargetID                    string  `json:"target_id" validate:"required"`
	PriceBeforeDiscount         float64 `json:"price_before_discount" validate:"required"`
	TotalDiscount               float64 `json:"total_discount" validate:"required"`
	PriceAfterDiscount          float64 `json:"price_after_discount" validate:"required"`
	AuthUserID                  string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateTransactionDiscountApplied struct {
	ID                  string  `json:"id" validate:"required" swaggerignore:"true"`
	TargetID            string  `json:"target_id" validate:"required"`
	PriceBeforeDiscount float64 `json:"price_before_discount" validate:"required"`
	TotalDiscount       float64 `json:"total_discount" validate:"required"`
	PriceAfterDiscount  float64 `json:"price_after_discount" validate:"required"`
	AuthUserID          string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}
