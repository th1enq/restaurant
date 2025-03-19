package employee

import (
	"log"
	"restaurant/drinking"
	"restaurant/food"
	"restaurant/order"
	"sync"
)

type Waiter struct {
	Employee
}

func (c *Waiter) Work(announcement <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range announcement {
		value, _ := val.(order.Order)
		foodName, isFood := value.Things.(food.IFood)
		drinkingName, _ := value.Things.(drinking.IDrinking)
		if isFood {
			log.Printf("Waiter %d serving %v for %v\n", c.ID, foodName.GetFoodName(), value.NameCustomer)
		} else {
			log.Printf("Waiter %d serving %v for %v\n", c.ID, drinkingName.GetDrinkingName(), value.NameCustomer)
		}
	}
}

func (c *Waiter) SetStatus(status string) {
	c.Status = status
}

func NewWaiter(id int) *Waiter {
	return &Waiter{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
