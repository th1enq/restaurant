package employee

import (
	"fmt"
	"log"
	"math/rand"
	"restaurant/drinking"
	"restaurant/order"
	"sync"
	"time"
)

type Bartender struct {
	Employee
}

func (b *Bartender) Work(readyDrinking chan<- interface{}, wg *sync.WaitGroup, thingsOrder order.Order) {
	defer func() {
		wg.Done()
		b.Mu.Lock()
		b.Status = RELAX
		b.Mu.Unlock()
	}()
	things, err := drinking.GetDrinking(thingsOrder.Things.(drinking.IDrinking).GetDrinkingName())
	if err != nil {
		panic(err)
	}
	recipe := things.GetDrinkingStep()

	for _, step := range recipe {
		b.Mu.Lock()
		b.Status = fmt.Sprintf("Bartender %d %s", b.ID, step)
		b.Mu.Unlock()
		log.Println(b.Status)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	readyDrinking <- thingsOrder
}

func newBartender(id int) *Bartender {
	return &Bartender{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
