package arrays

func replaceElements(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
	}

	var maxNum, maxNumIndex int

	for i := 0; i < len(arr)-1; i++ {
		if i < maxNumIndex {
			arr[i] = maxNum
			continue
		}

		maxNum = 0
		maxNumIndex = 0

		for j, num := range arr[i+1:] {
			if num > maxNum {
				maxNum = num
				maxNumIndex = j
			}
		}

		arr[i] = maxNum
	}

	arr[len(arr)-1] = -1

	return arr
}
