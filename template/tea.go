package template

type Tea struct{}

func (t *Tea) PrepareRecipe() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

func (t *Tea) BoilWater() {
	println("Boiling water for tea...")
}

func (t *Tea) Brew() {
	println("Steeping the tea...")
}

func (t *Tea) PourInCup() {
	println("Pouring tea into cup...")
}

func (t *Tea) AddCondiments() {
	println("Adding lemon to tea...")
}
