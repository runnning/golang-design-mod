package visitor

type Square struct {
	side float64
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}
