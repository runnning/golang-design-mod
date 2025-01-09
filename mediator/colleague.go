package mediator

type Colleague interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}
