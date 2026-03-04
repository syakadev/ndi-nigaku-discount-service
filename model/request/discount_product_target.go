package reqmodel

type CreateDiscountProductTarget struct {
	DiscountID          string  `json:"discount_id" validate:"required"`
	TargetID            string  `json:"target_id" validate:"required"`
	ProductName         string  `json:"product_name"`
	MaxTotalQuota       int     `json:"max_total_quota"`
	PriceBeforeDiscount float64 `json:"price_before_discount"`
	TotalDiscount       float64 `json:"total_discount"`
	PriceAfterDiscount  float64 `json:"price_after_discount"`
	IsActive            bool    `json:"is_active" validate:"required"`
	AuthUserID          string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateDiscountProductTarget struct {
	ID                  string  `json:"id" validate:"required" swaggerignore:"true"`
	DiscountID          string  `json:"discount_id" validate:"required"`
	TargetID            string  `json:"target_id" validate:"required"`
	ProductName         string  `json:"product_name"`
	MaxTotalQuota       int     `json:"max_total_quota"`
	PriceBeforeDiscount float64 `json:"price_before_discount"`
	TotalDiscount       float64 `json:"total_discount"`
	PriceAfterDiscount  float64 `json:"price_after_discount"`
	IsActive            bool    `json:"is_active" validate:"required"`
	AuthUserID          string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type CreateProductDiscountApplied struct {
	DiscountProductTargetID string `json:"discount_product_target_id" validate:"required"`
	CustomerID              string `json:"customer_id"`
	CustomerName            string `json:"customer_name"`
	TransactionDate         string `json:"transaction_date"`
	AuthUserID              string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}
