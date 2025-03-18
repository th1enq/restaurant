package drinking

type Coffee struct {
	Drinking
}

func NewCoffee() *Coffee {
	return &Coffee{
		Drinking{
			status: PENDING,
			recipe: map[int]string{
				1: "Grind the coffee beans",
				2: "Boil water",
				3: "Pour the water over the coffee grounds",
				4: "Stir the coffee",
			},
		},
	}
}

func (b *Coffee) GetDrinkingName() string {
	return COFFEE
}
