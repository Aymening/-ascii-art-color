package handler

// IsHere checks if a given index exists in the provided slice of integers.
// It iterates through the slice and returns true if the index is found, otherwise false.
func IsHere(i int, result []int) bool {
	for y := 0; y < len(result); y++ {
		if i == result[y] {
			return true
		}
	} 
	return false
}
