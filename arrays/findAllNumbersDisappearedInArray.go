package arrays

import (
	"slices"
)

func findDisappearedNumbers(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}

	slices.Sort(nums)

	var res []int

	for i := range len(nums) {
		if _, ok := slices.BinarySearch(nums, i+1); !ok {
			res = append(res, i+1)
		}
	}

	return res
}
