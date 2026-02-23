package reqmodel

type ListRequest struct {
	Page   int    `json:"page" example:"1" validate:"required"`
	Size   int    `json:"size" example:"10" validate:"required"`
	Search string `json:"search"`
}
