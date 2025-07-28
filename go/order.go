package order_processor

type Item struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	id     string
	items  []Item
	total  float64
	status string
}

func NewOrder(id string) *Order {
	order := Order{
		id:     id,
		status: "Pending",
	}

	return &order
}

func (o *Order) Total() float64 {
	return o.total
}

func (o *Order) Id() string {
	return o.id
}

func (o *Order) Status() string {
	return o.status
}

func (o *Order) AddItem(item Item) {
	o.items = append(o.items, item)
	o.CalculateTotal()
}

func (o *Order) CalculateTotal() {
	total := 0.0
	for _, item := range o.items {
		total += item.Price * float64(item.Quantity)
	}
	o.total = total
}
