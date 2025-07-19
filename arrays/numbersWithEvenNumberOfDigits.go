package arrays

func findNumbers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var mainCounter int = 0

	for _, num := range nums {
		currentCounter := 0

		for num > 0 {
			num /= 10
			currentCounter++
		}

		if currentCounter != 0 && currentCounter%2 == 0 {
			mainCounter++
		}
	}

	return mainCounter
}
