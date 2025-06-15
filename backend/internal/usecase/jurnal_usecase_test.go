package usecase

import (
	"context"
	"errors"
	"studia/backend/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJurnalRepository struct {
	mock.Mock
}

func (m *MockJurnalRepository) Save(ctx context.Context, jurnal *domain.Jurnal) error {
	args := m.Called(ctx, jurnal)
	return args.Error(0)
}

func (m *MockJurnalRepository) FindByID(ctx context.Context, id *uuid.UUID) (*domain.Jurnal, error){
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Jurnal), args.Error(1)
}

func (m *MockJurnalRepository) FindAll(ctx context.Context) ([]domain.Jurnal, error) {
	args := m.Called(ctx) 
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Jurnal), args.Error(1)
}
func (m *MockJurnalRepository) Update(ctx context.Context, jurnal *domain.Jurnal) error {
	args := m.Called(ctx, jurnal)
	return args.Error(0)
}

func (m *MockJurnalRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestJurnalUsecase_Create(t *testing.T){
	mockRepo := new(MockJurnalRepository)
	jurnalInput := &domain.Jurnal{
		ID:  uuid.New(),
		Activity: "Mengerjakan tugas Go",
		Description: "Membuat unit test untuk usecase",
	}
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

	uc := NewJurnalUsecase(mockRepo, 2*time.Second)

	err := uc.Create(context.Background(), jurnalInput)

	assert.NoError(t, err)
	assert.Equal(t, domain.StatusPending, jurnalInput.Status)
	mockRepo.AssertExpectations(t)
}

func TestJurnalUsecase_FindAll_Success(t *testing.T) {
	mockRepo := new(MockJurnalRepository)
	mockJurnals := []domain.Jurnal{
		{Activity: "Jurnal 1"},
		{Activity: "Jurnal 2"},
	}
	// Program mock untuk mengembalikan data
	mockRepo.On("FindAll", mock.Anything).Return(mockJurnals, nil)

	uc := NewJurnalUsecase(mockRepo, 2*time.Second)
	result, err := uc.FindAll(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestJurnalUsecase_FindAll_Error(t *testing.T) {
	mockRepo := new(MockJurnalRepository)
	expectedError := errors.New("database findall error")
	// Program mock untuk mengembalikan error
	mockRepo.On("FindAll", mock.Anything).Return(nil, expectedError)

	uc := NewJurnalUsecase(mockRepo, 2*time.Second)
	result, err := uc.FindAll(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}


// --- Tes untuk Fungsi Update ---

func TestJurnalUsecase_Update_Success(t *testing.T) {
	mockRepo := new(MockJurnalRepository)
	jurnalToUpdate := &domain.Jurnal{
		ID:       uuid.New(),
		Activity: "Aktivitas yang sudah diupdate",
	}
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Jurnal")).Return(nil)

	uc := NewJurnalUsecase(mockRepo, 2*time.Second)
	err := uc.Update(context.Background(), jurnalToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// --- Tes untuk Fungsi Delete ---

func TestJurnalUsecase_Delete_Success(t *testing.T) {
	mockRepo := new(MockJurnalRepository)
	idToDelete, _ := uuid.NewRandom()
	mockRepo.On("Delete", mock.Anything, &idToDelete).Return(nil)

	uc := NewJurnalUsecase(mockRepo, 2*time.Second)
	err := uc.Delete(context.Background(), &idToDelete)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
