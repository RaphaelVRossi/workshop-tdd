package order_processor

import (
	"testing"
)

func TestNewOrderAndCalculateTotal(t *testing.T) {
	items := []Item{
		{ID: "p1", Name: "Computador", Price: 1000.0, Quantity: 1},
		{ID: "p2", Name: "Mouse", Price: 25.0, Quantity: 2},
	}

	order := NewOrder("ORD001", items)

	if order.ID != "ORD001" {
		t.Errorf("Expected Order ID ORD001, got %s", order.ID)
	}
	if order.Status != "Pending" {
		t.Errorf("Expected initial status Pending, got %s", order.Status)
	}

	expectedTotal := 1000.0*1 + 25.0*2
	if order.Total != expectedTotal {
		t.Errorf("Expected total %f, got %f", expectedTotal, order.Total)
	}

	// Test with empty items
	emptyOrder := NewOrder("ORD002", []Item{})
	if emptyOrder.Total != 0.0 {
		t.Errorf("Expected empty order total 0.0, got %f", emptyOrder.Total)
	}
}
