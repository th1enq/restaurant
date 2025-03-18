package customer

import (
	"restaurant/employee"
	"restaurant/order"
)

type Customer struct {
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name: name,
	}
}

func (c *Customer) Order(things string) {
	employee.GetManager().AddOrder(*order.NewOrder(things, c.name))
}
