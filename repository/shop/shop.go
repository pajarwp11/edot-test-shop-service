package shop

import (
	"shop-service/models/shop"

	"github.com/jmoiron/sqlx"
)

type ShopRepository struct {
	mysql *sqlx.DB
}

func NewShopRepository(mysql *sqlx.DB) *ShopRepository {
	return &ShopRepository{
		mysql: mysql,
	}
}

func (s *ShopRepository) Insert(shop *shop.RegisterRequest) error {
	_, err := s.mysql.Exec("INSERT INTO shops (name,user_id) VALUES (?,?,)", shop.Name, shop.UserId)
	return err
}
