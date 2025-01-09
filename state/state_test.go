package state

import "testing"

func TestState(t *testing.T) {
	light := &Light{}

	// 初始状态为关
	light.SetState(&OffState{})

	// 按钮操作
	light.PressButton()
	light.PressButton()
}
