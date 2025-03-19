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

func (c *Chef) Work(readyFood chan<- interface{}, wg *sync.WaitGroup, thingOrder order.Order) {
	defer func() {
		wg.Done()
		c.Mu.Lock()
		c.Status = RELAX
		c.Mu.Unlock()
	}()

	things, err := food.GetFood(thingOrder.Things.(food.IFood).GetFoodName())
	if err != nil {
		panic(err)
	}
	recipe := things.GetFoodStep()

	for i := 0; i < len(recipe); i++ {
		c.Mu.Lock()
		c.Status = fmt.Sprintf("Chef %d %s", c.ID, recipe[i])
		c.Mu.Unlock()
		log.Println(c.Status)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	readyFood <- thingOrder
}

func newChef(id int) *Chef {
	return &Chef{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
