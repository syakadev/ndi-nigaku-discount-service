package reqmodel

type CreateAppliedDiscount struct {
	DiscountTargetID    string  `json:"discount_target_id" validate:"required"`
	PriceBeforeDiscount float64 `json:"price_before_discount" validate:"required"`
	TotalDiscount       float64 `json:"total_discount" validate:"required"`
	PriceAfterDiscount  float64 `json:"price_after_discount" validate:"required"`
	AuthUserID          string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdateAppliedDiscount struct {
	ID                  string  `json:"id" validate:"required" swaggerignore:"true"`
	DiscountTargetID    string  `json:"discount_target_id"`
	PriceBeforeDiscount float64 `json:"price_before_discount"`
	TotalDiscount       float64 `json:"total_discount"`
	PriceAfterDiscount  float64 `json:"price_after_discount"`
	AuthUserID          string  `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}
