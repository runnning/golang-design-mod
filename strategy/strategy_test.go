package strategy

import "testing"

func TestStrategy(t *testing.T) {
	// 用户选择支付宝支付
	alipay := &Alipay{}

	paymentContext := &PaymentContext{}
	paymentContext.SetStrategy(alipay)
	paymentContext.ExecutePayment(100.0)

	// 用户选择微信支付
	wechatPay := &WeChatPay{}
	paymentContext.SetStrategy(wechatPay)
	paymentContext.ExecutePayment(200.0)

	// 用户选择银行卡支付
	bankPay := &BankPay{}
	paymentContext.SetStrategy(bankPay)
	paymentContext.ExecutePayment(300.0)
}
