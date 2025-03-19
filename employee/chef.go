package employee

import (
	"fmt"
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

func (c *Chef) Work(readyFood chan<- interface{}, wg *sync.WaitGroup, orderLists chan order.Order) {
	defer wg.Done()

	for orderItem := range orderLists {
		fmt.Println(c, orderItem)
		foodItem, _ := food.GetFood(orderItem.Item.(food.IFood).GetFoodName())
		c.SetStatus(WORKING)
		for _, step := range foodItem.GetRecipe() {
			c.SetStatus(fmt.Sprintf("Chef %d %s", c.ID, step))
			log.Println(c.Status)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		readyFood <- orderItem
		c.SetStatus(RELAX)
	}
}

func newChef(id int) *Chef {
	return &Chef{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
