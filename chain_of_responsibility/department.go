package chain_of_responsibility

type Department interface {
	execute(*Patient)
	setNext(Department)
}
