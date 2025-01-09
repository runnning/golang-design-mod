package mediator

import "testing"

func TestMediator(t *testing.T) {
	chatRoom := &ChatRoom{}
	user1 := &User{
		name:     "test1",
		mediator: chatRoom,
	}
	user2 := &User{
		name:     "test2",
		mediator: chatRoom,
	}
	user3 := &User{
		name:     "test3",
		mediator: chatRoom,
	}
	chatRoom.RegisterUser(user1)
	chatRoom.RegisterUser(user2)
	chatRoom.RegisterUser(user3)

	user1.SendMessage("Hello, everyone!")
	user2.SendMessage("Hi, Alice!")
}
