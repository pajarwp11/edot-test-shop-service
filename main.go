package main

import (
	"fmt"
	"log"
	"net/http"
	"shop-service/db/mysql"
	shopHandler "shop-service/handler/shop"
	"shop-service/middleware"
	shopRepo "shop-service/repository/shop"
	shopUsecase "shop-service/usecase/shop"

	"github.com/gorilla/mux"
)

func main() {
	mysql.Connect()
	router := mux.NewRouter()
	shopRepository := shopRepo.NewShopRepository(mysql.MySQL)
	shopUsecase := shopUsecase.NewShopUsecase(shopRepository)
	shopHandler := shopHandler.NewShopHandler(shopUsecase)
	router.Handle("/shop/register", middleware.JWTMiddleware(http.HandlerFunc(shopHandler.Register))).Methods(http.MethodPost)

	fmt.Println("server is running")
	err := http.ListenAndServe(":8002", router)
	if err != nil {
		log.Fatal(err)
	}
}
