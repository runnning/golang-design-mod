package factory

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{Gun{
		name:  "ak47 gun",
		power: 4,
	}}
}
