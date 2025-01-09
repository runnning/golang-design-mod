package abstract_factory

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{Shoe{
		logo: "nike",
		size: 14,
	}}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{Shirt{
		logo: "nike",
		size: 14,
	}}
}
