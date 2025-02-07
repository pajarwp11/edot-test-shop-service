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

func (m *MockShopRepository) GetById(id int) (*shop.Shop, error) {
	args := m.Called(id)
	return args.Get(0).(*shop.Shop), args.Error(1)
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

func TestGetById_Success(t *testing.T) {
	mockRepo := new(MockShopRepository)
	shopUsecase := NewShopUsecase(mockRepo)

	mockShop := &shop.Shop{
		Id:   1,
		Name: "Shop A",
	}

	// Expect GetById to be called with ID 1 and return mockShop
	mockRepo.On("GetById", 1).Return(mockShop, nil)

	result, err := shopUsecase.GetById(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockShop, result)
	mockRepo.AssertExpectations(t)
}

func TestGetById_NotFound(t *testing.T) {
	mockRepo := new(MockShopRepository)
	shopUsecase := NewShopUsecase(mockRepo)

	// Simulate "shop not found" error
	mockRepo.On("GetById", 2).Return((*shop.Shop)(nil), errors.New("shop not found"))

	result, err := shopUsecase.GetById(2)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "shop not found", err.Error())
	mockRepo.AssertExpectations(t)
}
