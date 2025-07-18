package arrays

func findMaxConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, current := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			current++
		} else {
			if current > max {
				max = current
			}

			current = 1
		}
	}

	if current > max {
		max = current
	}

	return max
}

func findMaxConsecutiveOne(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max, current int

	for _, num := range nums {
		if num == 1 {
			current++
			if current > max {
				max = current
			}
		} else {
			current = 0
		}
	}

	return max
}
