package flyweight

import "fmt"

type CounterTerrorist struct {
	task   string
	weapon string
}

func (ct *CounterTerrorist) Task() {
	fmt.Printf("Counter Terrorist with weapon %s | Task: %s\n", ct.weapon, ct.task)
}
