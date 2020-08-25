package slice

// MapFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func MapFloat64(a []float64, f func(v float64) float64) []float64 {
	var b []float64
	for _, v := range a {
		b = append(b, f(v))
	}
	return b
}

// MapInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func MapInt(a []int, f func(v int) int) []int {
	var b []int
	for _, v := range a {
		b = append(b, f(v))
	}
	return b
}
