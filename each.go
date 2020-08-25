package slice

// EachFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func EachFloat64(a []float64, f func(v float64)) {
	for _, v := range a {
		f(v)
	}
}

// EachInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func EachInt(a []int, f func(v int)) {
	for _, v := range a {
		f(v)
	}
}
