package state

type State interface {
	PressButton(light *Light)
}
