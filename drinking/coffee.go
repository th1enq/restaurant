package drinking

type Coffee struct {
	Drinking
}

func NewCoffee() *Coffee {
	return &Coffee{
		Drinking{
			status: PENDING,
			recipe: []string{
				"Grind the coffee beans",
				"Boil water",
				"Pour the water over the coffee grounds",
				"Stir the coffee",
			},
		},
	}
}

func (b *Coffee) GetDrinkingName() string {
	return COFFEE
}
