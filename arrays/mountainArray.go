package arrays

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var upStep, downStep int64

	for i := 0; i < len(arr)-1; i++ {

		switch {
		case arr[i] == arr[i+1]:
			return false
		case arr[i] < arr[i+1] && downStep < 1:
			upStep++
		case arr[i] > arr[i+1] && upStep > 0:
			downStep++
		default:
			return false
		}
	}

	return upStep > 0 && downStep > 0
}
