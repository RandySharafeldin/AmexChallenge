package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RandySharafeldin/AmexChallenge/models"
)

type Orders struct {
	OrderService models.OrderService
}

func (o Orders) ProcessOrder(w http.ResponseWriter, r *http.Request) {

	order := models.Order{}
	apples, err := strconv.Atoi(r.FormValue("apples"))
	if err != nil {
		http.Error(w, "Invalid entry for apples", http.StatusBadRequest)
		return
	}
	oranges, err := strconv.Atoi(r.FormValue("oranges"))
	if err != nil {
		http.Error(w, "Invalid entry for oranges", http.StatusBadRequest)
		return
	}
	order.Apples = apples
	order.Oranges = oranges
	o.OrderService.MakeOrder(&order)
	w.WriteHeader(http.StatusOK)
	respBytes, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(respBytes)
}
