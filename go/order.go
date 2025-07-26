package order_processor

type Item struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	ID     string
	Items  []Item
	Total  float64
	Status string
}

func NewOrder(id string, items []Item) Order {
	order := Order{
		ID:     id,
		Items:  items,
		Status: "Pending",
	}

	order.CalculateTotal()

	return order
}

func (o *Order) CalculateTotal() {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.Total = total
}
