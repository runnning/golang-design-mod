package facade

import "fmt"

type CPU struct {
}

func (c *CPU) Start() {
	fmt.Println("CPU is starting")
}

func (c *CPU) Shutdown() {
	fmt.Println("CPU is shutting down")
}
