package template

type Coffee struct{}

func (c *Coffee) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

func (c *Coffee) BoilWater() {
	println("Boiling water for coffee...")
}

func (c *Coffee) Brew() {
	println("Brewing the coffee...")
}

func (c *Coffee) PourInCup() {
	println("Pouring coffee into cup...")
}

func (c *Coffee) AddCondiments() {
	println("Adding sugar and milk t")
}
