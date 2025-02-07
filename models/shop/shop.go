package shop

type Shop struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	UserId int    `db:"user_id"`
}

type RegisterRequest struct {
	Name   string `json:"name" validate:"required"`
	UserId int    `json:"user_id" validate:"required"`
}
