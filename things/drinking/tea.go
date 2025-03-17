package drinking

import "restaurant/helper"

type Tea struct {
	Drinking
}

func NewTea() *Tea {
	return &Tea{Drinking{status: PENDING}}
}

func (b *Tea) GetName() string {
	return TEA
}

func (b *Tea) GetStatus() string {
	return b.status
}

func (b *Tea) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
