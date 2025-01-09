package visitor

import "fmt"

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := 3.14 * c.radius * c.radius
	fmt.Printf("Area of Circle: %.2f\n", area)
}

func (a *AreaCalculator) VisitSquare(s *Square) {
	area := s.side * s.side
	fmt.Printf("Area of Square: %.2f\n", area)
}
