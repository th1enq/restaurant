package food

import "restaurant/helper"

type Pizza struct {
	Food
}

func NewPizza() *Pizza {
	return &Pizza{Food{status: PENDING}}
}

func (b *Pizza) GetName() string {
	return PIZZA
}

func (b *Pizza) GetStatus() string {
	return b.status
}

func (b *Pizza) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
