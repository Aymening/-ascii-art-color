package handler

// SearchSub searches for all occurrences of a substring within a given string.
// It returns a slice containing the indices of the substring occurrences.
func SearchSub(str, strsub string) []int {
	var result []int
	// Return an empty slice if the substring is empty
	if strsub == "" {
		return result
	}
	// Iterate through the string to find substring occurrences
	for i := 0; i <= len(str)-len(strsub); i++ {
		// If a substring occurrence is found, add its indices to the result slice
		if str[i:i+len(strsub)] == strsub {
			for j := 0; j < len(strsub); j++ {
				result = append(result, i+j)
			}
			// Move to the next character after the substring
			i += len(strsub) - 1
		}
	}
	return result
}
