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

func NewJurnalHandler(router *gin.RouterGroup, jurnalUsecase usecase.JurnalUsecase) {
	handler :=  &JurnalHandler{
		jurnalUsecase: jurnalUsecase,
	}

	publicJurnal := router.Group("journals")
	{
		publicJurnal.POST("/", handler.Create)
		publicJurnal.GET("/", handler.FindAll)
		publicJurnal.GET("/:id", handler.FindByID)
		publicJurnal.PUT("/:id", handler.Update)
		publicJurnal.DELETE("/:id", handler.Delete)
	}

}


func (h *JurnalHandler) Create(c *gin.Context){
	var req CreateJurnalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	jurnal := &domain.Jurnal{
		ID: uuid.New(),
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
		ID:          id,
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