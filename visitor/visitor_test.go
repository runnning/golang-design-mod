package visitor

import "testing"

func TestVisitor(t *testing.T) {
	shapes := []Shape{
		&Circle{radius: 5},
		&Square{side: 4},
	}

	areaCalculator := &AreaCalculator{}
	perimeterCalculator := &PerimeterCalculator{}

	for _, shape := range shapes {
		shape.Accept(areaCalculator)
		shape.Accept(perimeterCalculator)
	}
}
