package memento

type Originator struct {
	text string
}

func (o *Originator) SetText(text string) {
	o.text = text
}

func (o *Originator) GetText() string {
	return o.text
}

func (o *Originator) SaveTextToMemento() *Memento {
	return CreateMemento(o.text)
}

func (o *Originator) RestoreStateFromMemento(memento *Memento) {
	o.text = memento.GetText()
}
