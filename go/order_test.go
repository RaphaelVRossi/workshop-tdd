package order_processor

import (
	"testing"
)

func TestNewOrder(t *testing.T) {
	order := NewOrder("ORD001")

	if order.ID != "ORD001" {
		t.Errorf("Expected Order ID ORD001, got %s", order.ID)
	}
	if order.Status != "Pending" {
		t.Errorf("Expected initial status Pending, got %s", order.Status)
	}

	order2 := NewOrder("ORD002")
	if order.ID == order2.ID {
		t.Error("Expected new order with different ID")
	}
}

func TestAddItem(t *testing.T) {
	order := NewOrder("ORD002")

	t.Run("Add one item", func(t *testing.T) {
		item := Item{ID: "item1", Name: "Computer", Price: 1000.0, Quantity: 1}

		order.AddItem(item)

		expectedTotal := 1000.0 * 1
		if order.Total != expectedTotal {
			t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total)
		}
	})

	t.Run("Add some items", func(t *testing.T) {
		order.AddItem(Item{ID: "item2", Name: "Mouse", Price: 30.0, Quantity: 2})
		order.AddItem(Item{ID: "item3", Name: "Keyboard", Price: 15.0, Quantity: 1})

		expectedTotal := 1000.0*1 + 30.0*2 + 15.0*1 // 1075
		if order.Total != expectedTotal {
			t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total)
		}
	})
}

func TestCalculateTotal(t *testing.T) {
	order := NewOrder("ORD003")
	computerItem := Item{ID: "item1", Name: "Computer", Price: 1000.0, Quantity: 1}
	mouseItem := Item{ID: "item2", Name: "Mouse", Price: 30.0, Quantity: 1}

	t.Run("Calculate total with one item", func(t *testing.T) {
		order.Items = []Item{computerItem}
		order.CalculateTotal()

		expectedTotal := 1000.0 * 1
		if order.Total != expectedTotal {
			t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total)
		}
	})

	t.Run("Calculate total with some items", func(t *testing.T) {
		order.Items = []Item{computerItem, computerItem, mouseItem, mouseItem}
		order.CalculateTotal()

		expectedTotal := 1000.0*2 + 30.0*2
		if order.Total != expectedTotal {
			t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total)
		}
	})

	t.Run("Calculate total without items", func(t *testing.T) {
		emptyOrder := NewOrder("ORD004")
		if emptyOrder.Total != 0.0 {
			t.Errorf("Expected empty order total 0.0, got %f", emptyOrder.Total)
		}
	})
}
