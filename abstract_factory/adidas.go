package abstract_factory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShone{Shoe{
		logo: "adidas",
		size: 14,
	}}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{Shirt{
		logo: "adidas",
		size: 14,
	}}
}
