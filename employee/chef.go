package employee

import (
	"fmt"
	"log"
	"math/rand"
	"restaurant/food"
	"sync"
	"time"
)

type Chef struct {
	Employee
}

func (c *Chef) Work(readyFood chan<- interface{}, wg *sync.WaitGroup, foodName string) {
	defer func() {
		wg.Done()
		c.Mu.Lock()
		c.Status = RELAX
		c.Mu.Unlock()
	}()

	c.Mu.Lock()
	c.Status = "Working"
	c.Mu.Unlock()
	things, err := food.GetFood(foodName)
	if err != nil {
		panic(err)
	}
	recipe := things.GetFoodStep()

	for i := 0; i < len(recipe); i++ {
		c.Status = fmt.Sprintf("Chef %d %s", c.ID, recipe[i])
		log.Println(c.Status)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	readyFood <- things
}

func newChef(id int) *Chef {
	return &Chef{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
