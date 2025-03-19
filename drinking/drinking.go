package drinking

import (
	"restaurant/helper"
)

const (
	COFFEE = "coffee"
	JUICE  = "juice"
	TEA    = "tea"
	NONE   = "none"
)

type Drinking struct {
	Recipe []string
}

type IDrinking interface {
	GetDrinkingName() string
	GetRecipe() []string
}

func (b *Drinking) GetRecipe() []string {
	return b.Recipe
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
