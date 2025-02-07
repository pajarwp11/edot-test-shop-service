package shop

import (
	"errors"
	"shop-service/models/shop"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockShopRepository struct {
	mock.Mock
}

func (m *MockShopRepository) Insert(shop *shop.RegisterRequest) error {
	args := m.Called(shop)
	return args.Error(0)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockShopRepository)
	shopUsecase := NewShopUsecase(mockRepo)

	mockRequest := &shop.RegisterRequest{
		Name:    "Shop A",
		Address: "123 Main St",
	}

	// Expect Insert to be called with mockRequest and return nil (success)
	mockRepo.On("Insert", mockRequest).Return(nil)

	err := shopUsecase.Register(mockRequest)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_Fail(t *testing.T) {
	mockRepo := new(MockShopRepository)
	shopUsecase := NewShopUsecase(mockRepo)

	mockRequest := &shop.RegisterRequest{
		Name:    "Shop B",
		Address: "456 Another St",
	}

	// Simulate a database error
	mockRepo.On("Insert", mockRequest).Return(errors.New("database error"))

	err := shopUsecase.Register(mockRequest)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}
