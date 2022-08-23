package main

import (
	"net/http"

	"github.com/RandySharafeldin/AmexChallenge/controllers"
	"github.com/RandySharafeldin/AmexChallenge/models"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter();

	orderService := models.OrderService{}

	ordersC := controllers.Orders{OrderService: orderService}

	router.Post("/order", ordersC.ProcessOrder)

	http.ListenAndServe(":8080", router)
}