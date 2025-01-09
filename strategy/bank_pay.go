package strategy

type BankPay struct{}

func (b *BankPay) Pay(amount float64) {
	println("使用银行卡支付:", amount)
}
