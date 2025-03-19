package drinking

type Juice struct {
	Drinking
}

func NewJuice() *Juice {
	return &Juice{
		Drinking{
			Recipe: []string{
				"Peel the fruit",
				"Cut the fruit",
				"Blend the fruit",
			},
		},
	}
}

func (b *Juice) GetDrinkingName() string {
	return JUICE
}
