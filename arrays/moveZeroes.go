package arrays

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	k := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != 0 {
			k++
		}

		if nums[i-1] == 0 && nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}
