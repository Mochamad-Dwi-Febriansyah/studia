package http

import (
	"studia/backend/internal/usecase"
	"studia/backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(router *gin.RouterGroup, categoryUsecase usecase.CategoryUsecase) {
	handler := &CategoryHandler{
		categoryUsecase: categoryUsecase,
	}

	publicCategory := router.Group("categories")
	{
		publicCategory.POST("/", handler.Create)
		publicCategory.GET("/", handler.FindAll)
		publicCategory.GET("/:id", handler.FindByID)
		publicCategory.PUT("/:id", handler.Update)
		publicCategory.DELETE("/:id", handler.Delete)
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category := &domain.Category{
		ID:          uuid.New(),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		StatusView:  domain.StatusDraft,
	}

	if err := h.categoryUsecase.Create(c.Request.Context(), category); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	categories, err := h.categoryUsecase.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, categories)
}

func (h *CategoryHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.categoryUsecase.FindByID(c.Request.Context(), &categoryID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(200, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	category := &domain.Category{
		ID:          categoryID,
		Name:        *req.Name,
		Slug:        *req.Slug,
		Description: *req.Description,
		StatusView:  domain.StatusView(*req.StatusView),
	}

	if err := h.categoryUsecase.Update(c.Request.Context(), category); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.categoryUsecase.Delete(c.Request.Context(), &categoryID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}