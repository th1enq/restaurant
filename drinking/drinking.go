package drinking

import (
	"restaurant/helper"
)

const (
	PENDING = "pending"
	MAKING  = "making"
	SERVED  = "served"
)

const (
	COFFEE = "coffee"
	JUICE  = "juice"
	TEA    = "tea"
	NONE   = "none"
)

type Drinking struct {
	status string
	recipe []string
}

type IDrinking interface {
	GetDrinkingName() string
	GetDrinkingStatus() string
	SetDrinkingStatus(status string) error
	GetDrinkingStep() []string
}

func (b *Drinking) GetDrinkingStatus() string {
	return b.status
}

func (b *Drinking) SetDrinkingStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}

func (b *Drinking) GetDrinkingStep() []string {
	return b.recipe
}

func GetDrinking(name string) (IDrinking, error) {
	switch name {
	case COFFEE:
		return NewCoffee(), nil
	case JUICE:
		return NewJuice(), nil
	case TEA:
		return NewTea(), nil
	default:
		return nil, helper.ErrInvalidFood
	}
}
