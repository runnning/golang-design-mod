package template

import "testing"

func TestTemplate(t *testing.T) {
	// 客户端可以选择制作茶或咖啡
	var tea Beverage = &Tea{}
	tea.PrepareRecipe() // 输出制作茶的过程

	println("----")

	var coffee Beverage = &Coffee{}
	coffee.PrepareRecipe() // 输出制作咖啡的过程
}
