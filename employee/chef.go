package employee

import (
	"log"
	"math/rand"
	"restaurant/food"
	"restaurant/order"
	"sync"
	"time"
)

type Chef struct {
	Employee
}

func (c *Chef) Work(readyFood chan<- interface{}, wg *sync.WaitGroup, orderLists chan order.Order, workHistory *sync.Map) {
	for orderItem := range orderLists {
		<-c.Ready
		foodItem, _ := food.GetFood(orderItem.Item.(food.IFood).GetFoodName())
		for _, step := range foodItem.GetRecipe() {
			log.Printf("Chef %d %s", c.ID, step)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		c.Ready <- READY
		readyFood <- orderItem
		if val, ok := workHistory.Load(c); ok {
			workHistory.Store(c, val.(int)+1)
		} else {
			workHistory.Store(c, 1)
		}
		wg.Done()
	}
}

func NewChef(id int) *Chef {
	return &Chef{
		Employee: Employee{
			Ready: make(chan interface{}, 1),
			ID:    id,
		},
	}
}
