package arrays

import "math/rand"

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	return quickSort(nums)
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[rand.Intn(len(nums))]

	var less, equal, greater []int

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

func sortedSquaresFast(nums []int) []int {
	var left, right int = 0, len(nums) - 1

	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}

	return result
}
