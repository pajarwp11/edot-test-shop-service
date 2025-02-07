package shop

type Shop struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	UserId  int    `db:"user_id"`
}

type RegisterRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"sddress" validate:"required"`
	UserId  int    `json:"user_id" validate:"required"`
}
