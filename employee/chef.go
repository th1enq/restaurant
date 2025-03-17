package employee

type Chef struct {
	Employee
	readyFood chan interface{}
}

func (c *Chef) Work() {

}

func newChef() *Chef {
	return &Chef{
		Employee: Employee{
			Status: PENDING,
		},
		readyFood: make(chan interface{}),
	}
}

func (c *Chef) SetStatus(status string) {
	c.Status = status
}

func (c *Chef) GetReadyThing() chan interface{} {
	return c.readyFood
}
