package arrays

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var increasing, decreascing bool

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}

		if arr[i] < arr[i+1] && !increasing {
			continue
		}

		increasing = true

		if arr[i] > arr[i+1] && increasing && !decreascing {
			continue
		}
	}

	decreascing = true

	return increasing && decreascing
}
