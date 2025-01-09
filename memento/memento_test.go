package memento

import "testing"

func TestMemento(t *testing.T) {
	// 创建发起人（文本编辑器）
	originator := &Originator{}
	caretaker := &Caretaker{}

	// 用户输入文本
	originator.SetText("Hello")
	println("Current Text:", originator.GetText())

	// 保存当前状态
	caretaker.AddMemento(originator.SaveTextToMemento())

	// 用户修改文本
	originator.SetText("Hello, World!")
	println("Current Text:", originator.GetText())

	// 再次保存当前状态
	caretaker.AddMemento(originator.SaveTextToMemento())

	// 用户修改文本
	originator.SetText("Hello, Go!")
	println("Current Text:", originator.GetText())

	// 恢复到第一个状态
	originator.RestoreStateFromMemento(caretaker.GetMemento(0))
	println("Restored to first state:", originator.GetText())

	// 恢复到第二个状态
	originator.RestoreStateFromMemento(caretaker.GetMemento(1))
	println("Restored to second state:", originator.GetText())
}
