package sep_20_2022

import "fmt"

type AnotherTrignometryObject interface {
	area() float32
}

func CalcualteDoubleArea(shape AnotherTrignometryObject) {
	fmt.Println(2 * shape.area())
}
