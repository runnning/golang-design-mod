package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	computer := NewComputerFacade()

	// 开机
	computer.Start()

	fmt.Println("--- Working ---")

	// 关机
	computer.Shutdown()
}
