package state

import "fmt"

type OnState struct{}

func (s *OnState) PressButton(light *Light) {
	fmt.Println("Turning off the light.")
	light.SetState(&OffState{})
}
