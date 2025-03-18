package drinking

type Juice struct {
	Drinking
}

func NewJuice() *Juice {
	return &Juice{
		Drinking{
			status: PENDING,
			recipe: map[int]string{
				1: "Peel the fruit",
				2: "Cut the fruit",
				3: "Blend the fruit",
			},
		},
	}
}

func (b *Juice) GetDrinkingName() string {
	return JUICE
}
