package usecase

import (
	"context"
	"studia/backend/internal/domain"
	mocks "studia/backend/test/mock"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

 

func TestCategoryUsecase_Create(t *testing.T){
	mockRepo := new(mocks.CategoryRepository)
	categoryInput := &domain.Category{
		ID:          uuid.New(),
		Name:        "Programming",
		Slug:        "programming",
		Description: "All about programming languages and concepts.",
		StatusView:  domain.StatusDraft,
	}

	mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)

	err := uc.Create(context.Background(), categoryInput)

	assert.NoError(t, err)
	assert.Equal(t, categoryInput.Name, "Programming")
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_FindAll_Success(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	mockCategories := []domain.Category{}

	mockRepo.On("FindAll", mock.Anything).Return(mockCategories, nil)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	result, err := uc.FindAll(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)  
	assert.Equal(t, mockCategories, result)
	mockRepo.AssertExpectations(t)
}