package utils

type (
	StringSlice []string
)

func (slice StringSlice) Len() int {
	// Success
	return len(slice)
}

func (slice StringSlice) Contain(value string) bool {
	for _, item := range slice {
		if item == value {
			// Contain
			return true
		}
	}
	// Not contain
	return false
}

func (slice StringSlice) Reverse() StringSlice {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	// Success
	return slice
}
