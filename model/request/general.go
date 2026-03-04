package reqmodel

type ListRequest struct {
	Page   int    `json:"page" example:"1" `
	Size   int    `json:"size" example:"10"`
	Search string `json:"search"`
}
