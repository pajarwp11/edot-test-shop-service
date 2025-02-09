package shop

import (
	"database/sql"
	"errors"
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
	_, err := s.mysql.Exec("INSERT INTO shops (name,address,user_id) VALUES (?,?,?)", shop.Name, shop.Address, shop.UserId)
	return err
}

func (s *ShopRepository) GetById(id int) (*shop.Shop, error) {
	shop := shop.Shop{}
	err := s.mysql.Get(&shop, "SELECT id,name,address,user_id FROM shops WHERE id=?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("shop not exist")
		}
		return nil, err
	}
	return &shop, err
}
