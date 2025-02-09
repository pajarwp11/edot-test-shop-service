package shop

import (
	"encoding/json"
	"net/http"
	"shop-service/models/shop"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ShopUsecase interface {
	Register(shopRegister *shop.RegisterRequest) error
	GetById(id int) (*shop.Shop, error)
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

	userID := req.Header.Get("X-User-ID")
	request.UserId, _ = strconv.Atoi(userID)

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

func (s *ShopHandler) GetById(w http.ResponseWriter, req *http.Request) {
	response := Response{}
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "id is required"
		json.NewEncoder(w).Encode(response)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "id must be numeric"
		json.NewEncoder(w).Encode(response)
		return
	}

	shop, err := s.shopUsecase.GetById(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response.Message = "get shop success"
	response.Data = shop
	json.NewEncoder(w).Encode(response)
}
