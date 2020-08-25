package slice

// InsertFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func InsertFloat64(a []float64, e float64, i int) []float64 {
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = e
	return a
}

// InsertInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func InsertInt(a []int, e int, i int) []int {
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = e
	return a
}
