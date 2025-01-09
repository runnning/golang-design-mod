package state

import "fmt"

type OffState struct{}

func (s *OffState) PressButton(light *Light) {
	fmt.Println("Turning on the light.")
	light.SetState(&OnState{})
}
