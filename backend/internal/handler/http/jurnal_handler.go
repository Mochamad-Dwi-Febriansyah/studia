package http

import (
	"net/http"
	"studia/backend/internal/domain"
	"studia/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JurnalHandler struct {
	jurnalUsecase usecase.JurnalUsecase
}

func NewJurnalHandler(jurnalUsecase usecase.JurnalUsecase) *JurnalHandler {
	return &JurnalHandler{
		jurnalUsecase: jurnalUsecase,
	}
}

type CreateJurnalRequest struct {
	Activity string `json:"activity" binding:"required"`
	Description string 	`json:"description" binding:"required"`
}
 
type UpdateJurnalRequest struct {
    Activity    *string `json:"activity"`
    Description *string `json:"description"`
}

func (h *JurnalHandler) Create(c *gin.Context){
	var req CreateJurnalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	jurnal := &domain.Jurnal{
		Activity: req.Activity,
		Description: req.Description,
	}

	if err := h.jurnalUsecase.Create(c.Request.Context(), jurnal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, jurnal)
}

func (h *JurnalHandler) FindAll(c *gin.Context) {
	jurnals, err := h.jurnalUsecase.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jurnals)
}


func (h *JurnalHandler) FindByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID Format"})
		return
	}

	jurnal, err := h.jurnalUsecase.FindByID(c.Request.Context(), &id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Jurnal Not Found"})
		return
	}

	c.JSON(http.StatusOK, jurnal)
}

func (h *JurnalHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateJurnalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jurnal := &domain.Jurnal{
		ID:          &id,
		Activity:    *req.Activity,
		Description: *req.Description,
	}

	if err := h.jurnalUsecase.Update(c.Request.Context(), jurnal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jurnal)
}

func (h *JurnalHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.jurnalUsecase.Delete(c.Request.Context(), &id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jurnal deleted successfully"})
}