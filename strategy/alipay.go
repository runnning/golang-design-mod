package strategy

type Alipay struct {
}

func (a *Alipay) Pay(amount float64) {
	println("使用支付宝支付:", amount)
}
