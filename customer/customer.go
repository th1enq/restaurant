package customer

import (
	"log"
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
	log.Printf("%s orders %s", c.name, itemName)
	if foodOrder, err := food.GetFood(itemName); err == nil {
		employee.GetManager().AddFoodOrder(*order.NewOrder(foodOrder, c.name))
		return
	}
	if drinkOrder, err := drinking.GetDrinking(itemName); err == nil {
		employee.GetManager().AddDrinkOrder(*order.NewOrder(drinkOrder, c.name))
		return
	}
	panic(helper.ErrInvalidItem)
}
