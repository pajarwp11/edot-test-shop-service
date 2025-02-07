package shop

import (
	"encoding/json"
	"net/http"
	"shop-service/models/shop"

	"github.com/go-playground/validator/v10"
)

type ShopUsecase interface {
	Register(shopRegister *shop.RegisterRequest) error
}

type ShopHandler struct {
	shopUsecase ShopUsecase
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var validate = validator.New()

func NewShopHandler(shopUsecase ShopUsecase) *ShopHandler {
	return &ShopHandler{
		shopUsecase: shopUsecase,
	}
}

func (s *ShopHandler) Register(w http.ResponseWriter, req *http.Request) {
	request := shop.RegisterRequest{}
	response := Response{}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "invalid request body"
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := validate.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	err := s.shopUsecase.Register(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response.Message = "shop registered"
	json.NewEncoder(w).Encode(response)
}
