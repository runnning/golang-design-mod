package strategy

type WeChatPay struct {
}

func (w *WeChatPay) Pay(amount float64) {
	println("使用微信支付:", amount)
}
