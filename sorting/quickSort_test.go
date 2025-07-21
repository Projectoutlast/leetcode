package sorting

import (
	helper "leetcode/helpers"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{3, 1, 4, 1, 5, 9, 2},
			want: []int{1, 1, 2, 3, 4, 5, 9},
		},
		{
			nums: []int{10, 9, 8, 7, 6, 5},
			want: []int{5, 6, 7, 8, 9, 10},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{5, -1, 3, 0, -2, 4},
			want: []int{-2, -1, 0, 3, 4, 5},
		},
		{
			nums: []int{7, 7, 7, 7, 7},
			want: []int{7, 7, 7, 7, 7},
		},
		{
			nums: []int{2147483647, -2147483648, 0},
			want: []int{-2147483648, 0, 2147483647},
		},
		{
			nums: []int{-1, 2, -1, 2, 0},
			want: []int{-1, -1, 0, 2, 2},
		},
	}

	for _, test := range tests {
		if got := quickSort(test.nums); !helper.CompareSlices(got, test.want) {
			t.Errorf("quickSort(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
