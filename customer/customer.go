package customer

import (
	"restaurant/drinking"
	"restaurant/employee"
	"restaurant/food"
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
	foodOrder, err := food.GetFood(things)
	if err == nil {
		employee.GetManager().AddOrder(*order.NewOrder(foodOrder, c.name))
	} else {
		drinkingOrder, _ := drinking.GetDrinking(things)
		employee.GetManager().AddOrder(*order.NewOrder(drinkingOrder, c.name))
	}
}
