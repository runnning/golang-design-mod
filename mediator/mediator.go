package mediator

type Mediator interface {
	Notify(sender Colleague, message string)
}
