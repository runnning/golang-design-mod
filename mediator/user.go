package mediator

import "fmt"

type User struct {
	name     string
	mediator Mediator
}

func (u *User) SendMessage(message string) {
	u.mediator.Notify(u, message)
}

func (u *User) ReceiveMessage(message string) {
	fmt.Println(u.name+"received:", message)
}
