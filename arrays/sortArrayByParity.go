package arrays

func sortArrayByParity(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	evenPntr := 0
	oddPntr := len(nums) - 1

	for evenPntr < oddPntr {
		switch {
		case (nums[evenPntr]%2 == 0) && (nums[oddPntr]%2 != 0):
			evenPntr++
			oddPntr--
		case (nums[evenPntr]%2 != 0) && (nums[oddPntr]%2 != 0):
			oddPntr--
		case (nums[evenPntr]%2 != 0) && (nums[oddPntr]%2 == 0):
			nums[evenPntr], nums[oddPntr] = nums[oddPntr], nums[evenPntr]
			evenPntr++
			oddPntr--
		case (nums[evenPntr]%2 == 0) && (nums[oddPntr]%2 == 0):
			evenPntr++
		}
	}

	return nums
}
