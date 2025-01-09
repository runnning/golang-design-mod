package memento

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(memento *Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index < 0 || index >= len(c.mementos) {
		return nil
	}
	return c.mementos[index]
}
