package factory

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{Gun{
		name:  "musket",
		power: 1,
	}}
}
