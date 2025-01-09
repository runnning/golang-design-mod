package observer

type PhoneDisplay struct {
	temperature float64
}

func NewPhoneDisplay() *PhoneDisplay {
	return &PhoneDisplay{}
}

func (pd *PhoneDisplay) Update(temperature float64) {
	pd.temperature = temperature
	println("Phone Display updated to:", temperature)
}
