package sep_20_2022

import "testing"

func TestRectangle(t *testing.T) {
	rect := new(Rectangle)
	rect.Breadth = 12
	rect.Length = 10

	CalculateArea(rect)
	CalculatePerimeter(rect)
}

func TestCircle(t *testing.T) {
	circ := new(Circle)
	circ.Radius = 2

	CalculateArea(circ)
	CalculatePerimeter(circ)
}
