package arrays

import "slices"

func merge(nums1 []int, m int, nums2 []int, _ int) []int {
	slices.Sort(append(nums1[:m], nums2...))
	return nums1
}
