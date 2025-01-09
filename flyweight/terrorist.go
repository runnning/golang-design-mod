package flyweight

import "fmt"

type Terrorist struct {
	task   string
	weapon string
}

func (t *Terrorist) Task() {
	fmt.Printf("Terrorist with weapon %s | Task: %s\n", t.weapon, t.task)
}
