package employee

import (
	"fmt"
	"log"
	"math/rand"
	"restaurant/drinking"
	"restaurant/food"
	"sync"
	"time"
)

type Chef struct {
	Employee
}

func (c *Chef) Work(readyFood chan<- interface{}, wg *sync.WaitGroup, foodName string) {
	defer wg.Done()
	defer c.SetStatus(RELAX)

	things, err := food.GetFood(foodName)
	if err != nil {
		panic(err)
	}
	recipe := things.GetFoodStep()

	for _, step := range recipe {
		c.SetStatus(fmt.Sprintf("Chef %d %s", c.ID, step))
		log.Println(c.Status)

		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	readyFood <- things
	c.SetStatus(drinking.SERVED)
}

func newChef(id int) *Chef {
	return &Chef{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
