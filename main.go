package main

import (
	"net/http"

	"github.com/RandySharafeldin/AmexChallenge/controllers"
	"github.com/RandySharafeldin/AmexChallenge/models"
	"github.com/RandySharafeldin/AmexChallenge/utils"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	dbConfig := utils.NewDBConfig()
	db, err := utils.Open(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderService := models.OrderService{
		DB: db,
	}

	ordersC := controllers.Orders{OrderService: orderService}

	router.Post("/order", ordersC.ProcessOrder)
	router.Get("/order", ordersC.FindOrder)
	router.Get("/all", ordersC.AllOrders)

	http.ListenAndServe(":8080", router)
}
