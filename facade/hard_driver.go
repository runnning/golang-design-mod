package facade

import "fmt"

type HardDriver struct {
}

func (h *HardDriver) Read() {
	fmt.Println("HardDrive is reading")
}

func (h *HardDriver) Write() {
	fmt.Println("HardDrive is writing")
}
