package controllers

import (
	"database/sql"
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

func (o Orders) FindOrder(w http.ResponseWriter, r *http.Request) {

	urlValues := r.URL.Query()
	idString := urlValues.Get("id")

	if len(idString) == 0 {
		http.Error(w, "No id was sent in url", http.StatusBadRequest)
		return
	}
	id, err:= strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Id was not a valid integer", http.StatusBadRequest)
		return
	}
	if id < 0 {
		http.Error(w, "Id was not a valid integer", http.StatusBadRequest)
		return
	}

	order, err := o.OrderService.GetById(id)
	if err == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("Order %d does not exist", id), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "We had an issue finding this order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	respBytes, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(respBytes)
}

func (o Orders) AllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := o.OrderService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	respBytes, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(respBytes)
}
