package arrays

func checkIfExist(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	numberMap := make(map[int]struct{}, len(arr))

	for i := range arr {

		if i < len(arr)-1 {
			if arr[i] == 2*arr[i+1] {
				return true
			}

			if _, ok := numberMap[2*arr[i+1]]; ok {
				return true
			}
		}

		if arr[i]%2 == 0 {
			if _, ok := numberMap[arr[i]/2]; ok {
				return true
			}
		}

		numberMap[arr[i]] = struct{}{}
	}

	return false
}
