package arrays

import "testing"

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
