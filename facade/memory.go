package facade

import "fmt"

type Memory struct {
}

func (m *Memory) Load() {
	fmt.Println("Memory is loading")
}

func (m *Memory) Unload() {
	fmt.Println("Memory is unloading")
}
