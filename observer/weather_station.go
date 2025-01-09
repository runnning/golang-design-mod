package observer

type WeatherStation struct {
	observers   []Observer
	temperature float64
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{}
}

func (w *WeatherStation) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) RemoveObserver(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherStation) NotifyObservers() {
	for i := range w.observers {
		w.observers[i].Update(w.temperature)
	}
}

func (w *WeatherStation) SetTemperature(temperature float64) {
	w.temperature = temperature
	w.NotifyObservers()
}
