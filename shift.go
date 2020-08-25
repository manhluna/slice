package slice

// ShiftFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func ShiftFloat64(a []float64, e float64) []float64 {
	a = append(a, 0)
	copy(a[1:], a[0:])
	a[0] = e
	return a
}

// ShiftInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func ShiftInt(a []int, e int) []int {
	a = append(a, 0)
	copy(a[1:], a[0:])
	a[0] = e
	return a
}
