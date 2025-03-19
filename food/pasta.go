package food

type Pasta struct {
	Food
}

func NewPasta() *Pasta {
	return &Pasta{
		Food{
			Recipe: []string{
				"Boil the pasta",
				"Add the sauce",
				"Mix the pasta and sauce",
			},
		},
	}
}

func (b *Pasta) GetFoodName() string {
	return PASTA
}
