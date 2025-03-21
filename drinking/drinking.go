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

type DrinkingFactory func() IDrinking

var drinkingFactories = make(map[string]DrinkingFactory)

func RegisterDrinking(name string, factory DrinkingFactory) {
	drinkingFactories[name] = factory
}

func GetDrinking(name string) (IDrinking, error) {
	if factory, ok := drinkingFactories[name]; ok {
		return factory(), nil
	}
	return nil, helper.ErrInvalidDrinking
}
