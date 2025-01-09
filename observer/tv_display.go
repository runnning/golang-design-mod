package observer

type TVDisplay struct {
	temperature float64
}

func NewTVDisplay() *TVDisplay {
	return &TVDisplay{}
}

func (td *TVDisplay) Update(temperature float64) {
	td.temperature = temperature
	println("TV Display updated to:", temperature)
}
