package memento

type Memento struct {
	text string
}

func CreateMemento(text string) *Memento {
	return &Memento{text: text}
}

func (m *Memento) GetText() string {
	return m.text
}
