package employee

import (
	"fmt"
	"restaurant/drinking"
	"restaurant/food"
	"sync"
)

type Waiter struct {
	Employee
}

func (c *Waiter) Work(announcement <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range announcement {
		switch v := val.(type) {
		case food.IFood:
			fmt.Printf("Waiter %d served %s\n", c.ID, v.GetFoodName())
		case drinking.IDrinking:
			fmt.Printf("Waiter %d served %s\n", c.ID, v.GetDrinkingName())
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
