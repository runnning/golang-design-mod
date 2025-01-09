package state

type Light struct {
	state State
}

func (l *Light) SetState(state State) {
	l.state = state
}

func (l *Light) PressButton() {
	l.state.PressButton(l)
}
