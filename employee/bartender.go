package employee

import (
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

func (b *Bartender) Work(readyDrinking chan<- interface{}, wg *sync.WaitGroup, orderLists chan order.Order, workHistory *sync.Map) {

	for orderItem := range orderLists {
		<-b.Ready
		drinkItem, _ := drinking.GetDrinking(orderItem.Item.(drinking.IDrinking).GetDrinkingName())
		for _, step := range drinkItem.GetRecipe() {
			log.Printf("Bartender %d %s", b.ID, step)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		b.Ready <- READY
		readyDrinking <- orderItem
		if val, ok := workHistory.Load(b); ok {
			workHistory.Store(b, val.(int)+1)
		} else {
			workHistory.Store(b, 1)
		}
		wg.Done()
	}
}

func newBartender(id int) *Bartender {
	return &Bartender{
		Employee: Employee{
			Ready: make(chan interface{}, 1),
			ID:    id,
		},
	}
}
