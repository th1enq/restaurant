package drinking

type Tea struct {
	Drinking
}

func NewTea() *Tea {
	return &Tea{
		Drinking{
			status: PENDING,
			recipe: []string{
				"Boil water",
				"Steep the tea",
				"Pour tea into a cup",
			},
		},
	}
}

func (b *Tea) GetDrinkingName() string {
	return TEA
}
