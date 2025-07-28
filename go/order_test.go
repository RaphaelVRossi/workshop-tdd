package order_processor

import (
	"testing"
)

func TestNewOrder(t *testing.T) {
	order := NewOrder("ORD001")

	if order.Id() != "ORD001" {
		t.Errorf("Expected Order ID ORD001, got %s", order.Id())
	}
	if order.Status() != "Pending" {
		t.Errorf("Expected initial status Pending, got %s", order.Status())
	}

}

func TestCreateTwoOrder(t *testing.T) {
	order := NewOrder("ORD001")
	order2 := NewOrder("ORD002")
	if order.Id() == order2.Id() {
		t.Error("Expected new order with different ID")
	}
}

func TestAddOneItem(t *testing.T) {
	order := NewOrder("ORD002")

	item := Item{ID: "item1", Name: "Computer", Price: 1000.0, Quantity: 1}

	order.AddItem(item)

	expectedTotal := 1000.0 * 1
	if order.Total() != expectedTotal {
		t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total())
	}
}

func TestAddItem(t *testing.T) {
	order := NewOrder("ORD002")

	order.AddItem(Item{ID: "item2", Name: "Mouse", Price: 30.0, Quantity: 2})
	order.AddItem(Item{ID: "item3", Name: "Keyboard", Price: 15.0, Quantity: 1})

	expectedTotal := 30.0*2 + 15.0*1 // 75.0
	if order.Total() != expectedTotal {
		t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total())
	}
}
func TestCalculateTotalOneItem(t *testing.T) {
	order := NewOrder("ORD003")
	computerItem := Item{ID: "item1", Name: "Computer", Price: 1000.0, Quantity: 1}
	order.AddItem(computerItem)
	order.CalculateTotal()

	expectedTotal := 1000.0 * 1
	if order.Total() != expectedTotal {
		t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total())
	}
}

func TestCalculateTotalSomeItem(t *testing.T) {
	order := NewOrder("ORD003")
	computerItem := Item{ID: "item1", Name: "Computer", Price: 1000.0, Quantity: 1}
	mouseItem := Item{ID: "item2", Name: "Mouse", Price: 30.0, Quantity: 1}

	order.AddItem(computerItem)
	order.AddItem(computerItem)
	order.AddItem(mouseItem)
	order.AddItem(mouseItem)

	order.CalculateTotal()

	expectedTotal := 1000.0*2 + 30.0*2 // 1060.0
	if order.Total() != expectedTotal {
		t.Errorf("Expected total price %f, got %f", expectedTotal, order.Total())
	}
}

func TestCalculateTotalZeroItem(t *testing.T) {
	order := NewOrder("ORD004")

	if order.Total() != 0.0 {
		t.Errorf("Expected empty order total 0.0, got %f", order.Total())
	}
}

func TestRemoveOneItem(t *testing.T) {
	//TODO: implement
}

func TestRemoveAllItems(t *testing.T) {
	//TODO: implement
}

func TestApply0Discount(t *testing.T) {
	//TODO: implement
}

func TestApply100Discount(t *testing.T) {
	//TODO: implement
}

func TestApply50Discount(t *testing.T) {
	// TODO: implement
}

func TestCalculateDeliveryFee(t *testing.T) {
	// TODO: implement
}

func TestCalculateDeliveryFeeFar(t *testing.T) {
	// TODO: implement
}
