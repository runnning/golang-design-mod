package visitor

type Shape interface {
	Accept(visitor Visitor)
}
