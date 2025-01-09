package template

type Beverage interface {
	PrepareRecipe()

	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}
