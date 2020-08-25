package slice

// UnshiftFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func UnshiftFloat64(a []float64, e float64) []float64 {
	copy(a[0:], a[1:])
	return a[:len(a)-1]
}

// UnshiftInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func UnshiftInt(a []int, e int) []int {
	copy(a[0:], a[1:])
	return a[:len(a)-1]
}
