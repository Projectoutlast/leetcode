package arrays

func duplicateZeros(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
		}
	}

	return arr
}
