package observer

import "testing"

func TestObserver(t *testing.T) {
	// 创建具体的天气站（被观察者）
	weatherStation := NewWeatherStation()

	// 创建两个观察者（显示设备）
	phoneDisplay := NewPhoneDisplay()
	tvDisplay := NewTVDisplay()

	// 注册观察者
	weatherStation.RegisterObserver(phoneDisplay)
	weatherStation.RegisterObserver(tvDisplay)

	// 设置新的温度并通知观察者
	weatherStation.SetTemperature(25.3)

	// 移除一个观察者（例如：电视显示）
	weatherStation.RemoveObserver(tvDisplay)

	// 再次设置新的温度并通知观察者
	weatherStation.SetTemperature(30.1)
}
