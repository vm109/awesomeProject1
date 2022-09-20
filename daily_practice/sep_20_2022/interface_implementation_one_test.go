package sep_20_2022

import "testing"

func TestCalcualteDoubleAreaCircle(t *testing.T) {
	circ := new(Circle)
	circ.Radius = 4

	CalcualteDoubleArea(circ)
}
