package test

import (
	"testing"

	"github.com/RandySharafeldin/AmexChallenge/models"
)

func TestMakeOrder(t *testing.T) {
	test_order := models.Order{
		Apples:  4,
		Oranges: 4,
	}

	service := models.OrderService{}
	service.MakeOrder(&test_order)
	wantResult := float64(test_order.Apples)*0.6 + float64(test_order.Oranges)*0.25
	if test_order.Cost != wantResult {
		t.Errorf("got %f want %f", test_order.Cost, wantResult)
	}
}
