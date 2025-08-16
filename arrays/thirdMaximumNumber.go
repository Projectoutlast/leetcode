package arrays

import "slices"

func thirdMax(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return slices.Max(nums)
	}

	slices.Sort(nums)

	var step, resIdx int = 1, len(nums)-1

	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			step++
		}

		if step == 3 {
			resIdx = i
			break
		}
	}

	if step < 3 {
		return slices.Max(nums)
	}

	return nums[resIdx]
}
