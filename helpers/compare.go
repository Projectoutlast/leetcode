package helper

func CompareSlices(result, want []int) bool {
	if len(result) != len(want) {
		return false
	}

	for i, num := range result {
		if num != want[i] {
			return false
		}
	}

	return true
}