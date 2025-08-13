package sorting

func bubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	rightPointer := len(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < rightPointer-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		rightPointer--

	}

	return arr
}
