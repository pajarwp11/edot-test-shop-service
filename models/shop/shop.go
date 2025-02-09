package shop

type Shop struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
	UserId  int    `db:"user_id" json:"user_id"`
}

type RegisterRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	UserId  int
}
