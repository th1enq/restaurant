package employee

type Bartender struct {
	Employee
	readyDrinking chan interface{}
}

func (b *Bartender) Work() {

}

func newBartender() *Bartender {
	return &Bartender{
		Employee: Employee{
			Status: PENDING,
		},
		readyDrinking: make(chan interface{}),
	}
}

func (b *Bartender) SetStatus(status string) {
	b.Status = status
}

func (b *Bartender) GetReadyThing() chan interface{} {
	return b.readyDrinking
}
