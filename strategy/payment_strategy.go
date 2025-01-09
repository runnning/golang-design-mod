package strategy

type PaymentStrategy interface {
	Pay(amount float64)
}
