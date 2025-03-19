package order

type Order struct {
	Item         interface{}
	NameCustomer string
}

func NewOrder(itemOrder interface{}, name string) *Order {
	return &Order{
		Item:         itemOrder,
		NameCustomer: name,
	}
}
