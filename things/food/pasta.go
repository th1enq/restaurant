package food

import "restaurant/helper"

type Pasta struct {
	Food
}

func NewPasta() *Pasta {
	return &Pasta{Food{status: PENDING}}
}

func (b *Pasta) GetName() string {
	return PASTA
}

func (b *Pasta) GetStatus() string {
	return b.status
}

func (b *Pasta) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
