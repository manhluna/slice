package slice

// PushFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PushFloat64(a []float64, e float64) []float64 {
	return append(a, e)
}

// PushInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PushInt(a []int, e int) []int {
	return append(a, e)
}
