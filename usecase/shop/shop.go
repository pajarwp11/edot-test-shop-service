package shop

import (
	"shop-service/models/shop"
)

type ShopRepository interface {
	Insert(shop *shop.RegisterRequest) error
	GetById(id int) (*shop.Shop, error)
}

type ShopUsecase struct {
	shopRepo ShopRepository
}

func NewShopUsecase(shopRepo ShopRepository) *ShopUsecase {
	return &ShopUsecase{
		shopRepo: shopRepo,
	}
}

func (s *ShopUsecase) Register(shopRegister *shop.RegisterRequest) error {
	return s.shopRepo.Insert(shopRegister)
}

func (s *ShopUsecase) GetById(id int) (*shop.Shop, error) {
	return s.shopRepo.GetById(id)
}
