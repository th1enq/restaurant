package customer

import (
	"restaurant/drinking"
	"restaurant/employee"
	"restaurant/food"
	"restaurant/helper"
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

func (c *Customer) Order(itemName string) {
	if foodOrder, err := food.GetFood(itemName); err == nil {
		employee.GetManager().AddOrder(*order.NewOrder(foodOrder, c.name))
		return
	}
	if drinkOrder, err := drinking.GetDrinking(itemName); err == nil {
		employee.GetManager().AddOrder(*order.NewOrder(drinkOrder, c.name))
		return
	}
	panic(helper.ErrInvalidItem)
}
