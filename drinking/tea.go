package drinking

type Tea struct {
	Drinking
}

func NewTea() *Tea {
	return &Tea{
		Drinking{
			status: PENDING,
			recipe: map[int]string{
				1: "Boil water",
				2: "Steep the tea",
				3: "Pour tea into a cup",
			},
		},
	}
}

func (b *Tea) GetDrinkingName() string {
	return TEA
}
