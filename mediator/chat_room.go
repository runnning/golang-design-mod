package mediator

type ChatRoom struct {
	users []Colleague
}

func (c *ChatRoom) RegisterUser(user Colleague) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) Notify(sender Colleague, message string) {
	for _, user := range c.users {
		if user != sender {
			user.ReceiveMessage(message)
		}
	}
}
