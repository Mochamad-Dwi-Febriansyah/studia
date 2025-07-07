package http

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
	StatusView  string `json:"status_view" binding:"required,oneof=draft published"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
	StatusView  *string `json:"status_view" binding:"omitempty,oneof=draft published"`
}