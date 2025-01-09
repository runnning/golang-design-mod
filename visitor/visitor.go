package visitor

type Visitor interface {
	VisitCircle(c *Circle)
	VisitSquare(s *Square)
}
