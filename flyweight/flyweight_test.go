package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	factory := NewPlayerFactory()

	// 获取玩家对象
	player1 := factory.GetPlayer("T")
	player2 := factory.GetPlayer("CT")
	player3 := factory.GetPlayer("T")  // 复用已存在的对象
	player4 := factory.GetPlayer("CT") // 复用已存在的对象

	player1.Task()
	player2.Task()
	player3.Task()
	player4.Task()

	// 验证对象复用
	fmt.Printf("\nPlayer1 address: %p", player1)
	fmt.Printf("\nPlayer3 address: %p", player3)
	fmt.Printf("\nPlayer2 address: %p", player2)
	fmt.Printf("\nPlayer4 address: %p\n", player4)
}
