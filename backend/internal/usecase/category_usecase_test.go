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

func TestCategoryUsecase_FindAll_Error(t *testing.T){
	mockRepo := new(mocks.CategoryRepository)
	mockRepo.On("FindAll", mock.Anything).Return(nil, assert.AnError)
	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	_, err := uc.FindAll(context.Background())

	assert.Error(t, err) 
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Delete_Success(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryID := uuid.New()
	categoryIDPtr := &categoryID

	mockRepo.On("Delete", mock.Anything, categoryIDPtr).Return(nil)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	err := uc.Delete(context.Background(), categoryIDPtr)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Delete_Error(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryID := uuid.New()
	categoryIDPtr := &categoryID

	mockRepo.On("Delete", mock.Anything, categoryIDPtr).Return(assert.AnError)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	err := uc.Delete(context.Background(), categoryIDPtr)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Update_Success(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryToUpdate := &domain.Category{
		ID:          uuid.New(),
		Name:        "Updated Category",
		Slug:        "updated-category",
		Description: "Updated description",
		StatusView:  domain.StatusDraft,
	}

	mockRepo.On("Update", mock.Anything, categoryToUpdate).Return(nil)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	err := uc.Update(context.Background(), categoryToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_FindByID_Success(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryID := uuid.New()
	categoryIDPtr := &categoryID
	expectedCategory := &domain.Category{
		ID:          categoryID,
		Name:        "Test Category",
		Slug:        "test-category",
		Description: "This is a test category",
		StatusView:  domain.StatusDraft,
	}

	mockRepo.On("FindByID", mock.Anything, categoryIDPtr).Return(expectedCategory, nil)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	result, err := uc.FindByID(context.Background(), categoryIDPtr)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_FindByID_Error(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryID := uuid.New()
	categoryIDPtr := &categoryID

	mockRepo.On("FindByID", mock.Anything, categoryIDPtr).Return(nil, assert.AnError)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	result, err := uc.FindByID(context.Background(), categoryIDPtr)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Update_Error(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryToUpdate := &domain.Category{
		ID:          uuid.New(),
		Name:        "Category to Update",
		Slug:        "category-to-update",
		Description: "Description for category to update",
		StatusView:  domain.StatusDraft,
	}

	mockRepo.On("Update", mock.Anything, categoryToUpdate).Return(assert.AnError)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)
	err := uc.Update(context.Background(), categoryToUpdate)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Create_Error(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	categoryInput := &domain.Category{
		ID:          uuid.New(),
		Name:        "Error Category",
		Slug:        "error-category",
		Description: "This category will trigger an error",
		StatusView:  domain.StatusDraft,
	}

	mockRepo.On("Save", mock.Anything, mock.Anything).Return(assert.AnError)

	uc := NewCategoryUsecase(mockRepo, 2*time.Second)

	err := uc.Create(context.Background(), categoryInput)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}