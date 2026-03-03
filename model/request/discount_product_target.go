package reqmodel

type CreateDiscountProductTarget struct {
	DiscountID string `json:"discount_id" validate:"required"`
	TargetType string `json:"target_type" validate:"required"`
	TargetID string `json:"target_id" validate:"required"`
	MaxTotalQuota int `json:"max_total_quota" validate:"required"`
	IsActive   bool   `json:"is_active" validate:"required"`
	AuthUserID string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateDiscountProductTarget struct {
	ID string `json:"id" validate:"required" swaggerignore:"true"`
	DiscountID string `json:"discount_id" validate:"required"`
	TargetType string `json:"target_type" validate:"required"`
	TargetID string `json:"target_id" validate:"required"`
	MaxTotalQuota int `json:"max_total_quota" validate:"required"`
	IsActive   bool   `json:"is_active" validate:"required"`
	AuthUserID string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type CreateProductDiscountApplied struct {
	DiscountProductTargetID string `json:"discount_product_target_id" validate:"required"`
	CustomerID	string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	TransactionDate string `json:"transaction_date"`
	AuthUserID string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}