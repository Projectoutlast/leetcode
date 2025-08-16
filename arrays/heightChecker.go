package arrays

import "math/rand"

func heightChecker(heights []int) int {
	if len(heights) < 2 {
		return 0
	}

	expected := make([]int, len(heights))

	copy(expected, heights)

	expected = qSort(expected)

	var res int

	for i := range heights {
		if heights[i] != expected[i] {
			res++
		}
	}

	return res
}

func qSort(nums []int) []int {
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

	return append(append(qSort(less), equal...), qSort(greater)...)
}
