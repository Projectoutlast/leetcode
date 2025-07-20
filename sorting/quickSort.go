package sorting

import "math/rand"

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[rand.Intn(len(nums))]

	less := make([]int, 0, len(nums))
	equal := make([]int, 0, len(nums))
	greater := make([]int, 0, len(nums))

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		}

		if num == pivot {
			equal = append(equal, num)
		}

		if num > pivot {
			greater = append(greater, num)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}
