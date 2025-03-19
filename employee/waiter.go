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

func (w *Waiter) Work(announcement <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range announcement {
		orderItem, ok := val.(order.Order)
		if !ok {
			log.Printf("Waiter %d received an invalid order: %v\n", w.ID, val)
			continue
		}

		if foodItem, isFood := orderItem.Item.(food.IFood); isFood {
			log.Printf("Waiter %d serving food: %s for customer: %s\n", w.ID, foodItem.GetFoodName(), orderItem.NameCustomer)
		} else if drinkItem, isDrink := orderItem.Item.(drinking.IDrinking); isDrink {
			log.Printf("Waiter %d serving drink: %s for customer: %s\n", w.ID, drinkItem.GetDrinkingName(), orderItem.NameCustomer)
		} else {
			log.Printf("Waiter %d received an unknown item type in order: %v\n", w.ID, orderItem)
		}
	}
}

func (w *Waiter) SetStatus(status string) {
	w.Status = status
}

func NewWaiter(id int) *Waiter {
	return &Waiter{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
