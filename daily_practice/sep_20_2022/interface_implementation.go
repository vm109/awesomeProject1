package sep_20_2022

import (
	"fmt"
	"math"
)

type TrignometryObject interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	Length  float32
	Breadth float32
}
type Circle struct {
	Radius float32
}

func (r *Rectangle) area() float32 {
	return r.Length * r.Breadth
}

func (r *Rectangle) perimeter() float32 {
	return 2*r.Length + 2*r.Breadth
}

func (c *Circle) area() float32 {
	return float32((math.Pi) * math.Pow(float64(c.Radius), 2))
}

func (c *Circle) perimeter() float32 {
	return float32(math.Pi) * (c.Radius) * 2
}

func CalculateArea(shape TrignometryObject) {
	fmt.Println(shape.area())
}

func CalculatePerimeter(shape TrignometryObject) {
	fmt.Println(shape.perimeter())
}
