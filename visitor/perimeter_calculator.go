package visitor

import "fmt"

type PerimeterCalculator struct{}

func (p *PerimeterCalculator) VisitCircle(c *Circle) {
	perimeter := 2 * 3.14 * c.radius
	fmt.Printf("Perimeter of Circle: %.2f\n", perimeter)
}

func (p *PerimeterCalculator) VisitSquare(s *Square) {
	perimeter := 4 * s.side
	fmt.Printf("Perimeter of Square: %.2f\n", perimeter)
}
