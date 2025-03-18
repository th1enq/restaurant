package employee

import (
	"fmt"
	"log"
	"math/rand"
	"restaurant/drinking"
	"sync"
	"time"
)

type Bartender struct {
	Employee
}

func (b *Bartender) Work(readyDrinking chan<- interface{}, wg *sync.WaitGroup, drinkingName string) {
	defer wg.Done()
	defer b.SetStatus(RELAX)
	things, err := drinking.GetDrinking(drinkingName)
	if err != nil {
		panic(err)
	}
	recipe := things.GetDrinkingStep()

	for _, step := range recipe {
		b.SetStatus(fmt.Sprintf("Bartender %d %s", b.ID, step))
		log.Println(b.Status)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	readyDrinking <- things
	b.SetStatus(drinking.SERVED)
}

func newBartender(id int) *Bartender {
	return &Bartender{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
