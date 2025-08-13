package sorting

func choosingSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		minNum := arr[i]
		minNumIdx := i

		for j := i; j < len(arr); j++ {
			if arr[j] < minNum {
				minNum = arr[j]
				minNumIdx = j
			}
		}

		arr[i], arr[minNumIdx] = arr[minNumIdx], arr[i]
	}

	return arr
}
