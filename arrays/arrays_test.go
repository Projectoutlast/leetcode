package arrays

import (
	"slices"
	"testing"
)

func TestFindMaxConsecutive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			nums: []int{5, 6, 0, 6, 6, 6, 7, 7, 7, 7, 8, 8, 10, 23},
			want: 4,
		},
		{
			nums: []int{3, 0, 0, 0, 0, 0, 133, 133, 133, 133, 133, 133, 133, 133},
			want: 8,
		},
		{
			nums: []int{0},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := findMaxConsecutive(test.nums); got != test.want {
			t.Errorf("findMaxConsecutive(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestFindMaxConsecutiveOne(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 1, 1},
			want: 2,
		},
		{
			nums: []int{1, 0, 0, 1, 1, 1},
			want: 3,
		},
		{
			nums: []int{1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 8,
		},
		{
			nums: []int{0},
			want: 0,
		},
		{
			nums: []int{1},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := findMaxConsecutiveOne(test.nums); got != test.want {
			t.Errorf("findMaxConsecutiveOnes(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestFindNumbers(t *testing.T) {
	t.Parallel()

	tests := []struct{
		nums []int
		want int
	}{
		{
			nums: []int{12,345,2,6,7896},
			want: 2,
		},
		{
			nums: []int{555,901,482,1771},
			want: 1,
		},
		{
			nums: []int{10,0,3404,34958443},
			want: 3,
		},
	}

	for _, test := range tests {
		if got := findNumbers(test.nums); got != test.want {
			t.Errorf("findNumbers(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestSortedSquares(t *testing.T) {
	t.Parallel()

	tests := []struct{
		nums []int
		want []int
	}{
		{
			nums: []int{-4,-1,0,3,10},
			want: []int{0,1,9,16,100},
		},
		{
			nums: []int{-7,-3,2,3,11},
			want: []int{4,9,9,49,121},
		},
	}

	for _, test := range tests {
		if got := sortedSquares(test.nums); !slices.Equal(got, test.want) {
			t.Errorf("sortedSquares(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestSortedSquaresFast(t *testing.T) {
	t.Parallel()

	tests := []struct{
		nums []int
		want []int
	}{
		{
			nums: []int{-4,-1,0,3,10},
			want: []int{0,1,9,16,100},
		},
		{
			nums: []int{-7,-3,2,3,11},
			want: []int{4,9,9,49,121},
		},
	}

	for _, test := range tests {
		if got := sortedSquaresFast(test.nums); !slices.Equal(got, test.want) {
			t.Errorf("sortedSquaresFast(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
