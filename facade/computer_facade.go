package facade

import "fmt"

type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
	hd     *HardDriver
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:    &CPU{},
		memory: &Memory{},
		hd:     &HardDriver{},
	}
}

func (c *ComputerFacade) Start() {
	c.cpu.Start()
	c.memory.Load()
	c.hd.Read()
	fmt.Println("Computer is ready to use")
}

func (c *ComputerFacade) Shutdown() {
	c.cpu.Shutdown()
	c.memory.Unload()
	c.hd.Write()
	fmt.Println("Computer is shutting down")
}
