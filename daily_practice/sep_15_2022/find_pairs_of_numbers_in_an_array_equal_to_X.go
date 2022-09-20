package sep_15_2022

type Data struct {
	Arr        []int
	FinalPairs []*Pair
}

type Pair struct {
	first  int
	second int
}

/*
	Time Complexity -  O(n2)
	Reason - we are iterating through the array n * n times.
    As it is a loop inside a loop.
*/
func (d *Data) FindPairsOfNumbersInAnArrayWhoseProductIsValueX(X int) {
	pairs := make([]*Pair, 0)
	for i, val := range d.Arr {
		// this gives the quotient after division
		var quotient = float32(X) / float32(val)
		for j := 0; j < len(d.Arr) && j != i; j++ {
			if quotient == float32(d.Arr[j]) {
				pairs = append(pairs, &Pair{first: val, second: d.Arr[j]})
			}
		}
	}

	d.FinalPairs = pairs
}
