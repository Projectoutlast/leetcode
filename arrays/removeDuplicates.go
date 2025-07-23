package arrays

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
