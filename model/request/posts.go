package reqmodel

type CreatePost struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	IsActive   bool   `json:"is_active" validate:"required"`
	AuthUserID string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}

type UpdatePost struct {
	ID         string `json:"id" validate:"required" swaggerignore:"true"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsActive   bool   `json:"is_active"`
	AuthUserID string `json:"auth_user_id" validate:"required" swaggerignore:"true"`
}
