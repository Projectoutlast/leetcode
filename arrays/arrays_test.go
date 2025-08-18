package arrays

import (
	"slices"
	"testing"

	helper "leetcode/helpers"
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

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
		{
			nums: []int{10, 0, 3404, 34958443},
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

	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		if got := sortedSquares(test.nums); !helper.CompareSlices(got, test.want) {
			t.Errorf("sortedSquares(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestSortedSquaresFast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		if got := sortedSquaresFast(test.nums); !helper.CompareSlices(got, test.want) {
			t.Errorf("sortedSquaresFast(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestDuplicateZeros(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			arr:  []int{5, 0, 0, 7, 8, 10, 0, 0},
			want: []int{5, 0, 0, 0, 0, 7, 8, 10},
		},
	}

	for _, test := range tests {
		if got := duplicateZeros(test.arr); !helper.CompareSlices(got, test.want) {
			t.Errorf("duplicateZeros(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, test := range tests {
		if got := merge(test.nums1, test.m, test.nums2, test.n); !helper.CompareSlices(got, test.want) {
			t.Errorf("merge(%v, %v, %v, %v) = %v, want %v", test.nums1, test.m, test.nums2, test.n, got, test.want)
		}
	}
}

func TestRemoveElements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums      []int
		val       int
		wantSlice []int
		wantValue int
	}{
		{
			nums:      []int{3, 2, 2, 3},
			val:       3,
			wantSlice: []int{2, 2, 0, 0},
			wantValue: 2,
		},
		{
			nums:      []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:       2,
			wantSlice: []int{0, 1, 4, 0, 3, 0, 0, 0},
			wantValue: 5,
		},
		{
			nums:      []int{1},
			val:       1,
			wantSlice: []int{1},
			wantValue: 0,
		},
	}

	for _, test := range tests {
		if k := removeElement(test.nums, test.val); k != test.wantValue {
			t.Errorf("removeElement = %v, want %v", k, test.wantValue)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		{
			nums: []int{1, 1},
			want: 1,
		},
		{
			nums: []int{1, 1, 1},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := removeDuplicates(test.nums); got != test.want {
			t.Errorf("removeDuplicates(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestCheckIfExist(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
		{
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
	}

	for _, test := range tests {
		if got := checkIfExist(test.arr); got != test.want {
			t.Errorf("checkIfExist(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestValidMountainArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{2, 1},
			want: false,
		},
		{
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			arr:  []int{0, 1, 2, 3, 4, 3, 2, 5, 1, 0},
			want: false,
		},
		{
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
		{
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
	}

	for _, test := range tests {
		if got := validMountainArray(test.arr); got != test.want {
			t.Errorf("validMountainArray(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestReplaceElements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			arr:  []int{400},
			want: []int{-1},
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{5, 5, 5, 5, -1},
		},
	}

	for _, test := range tests {
		if got := replaceElements(test.arr); !helper.CompareSlices(got, test.want) {
			t.Errorf("replaceElements(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			arr:  []int{0},
			want: []int{0},
		},
		{
			arr:  []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			arr:  []int{1, 1, 0, 0, 3, 95, 0, 35, 45},
			want: []int{1, 1, 3, 95, 35, 45, 0, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.arr)
		if !slices.Equal(test.arr, test.want) {
			t.Errorf("moveZeroes = %v, want %v", test.arr, test.want)
		}
	}
}

func TestSortArrayByParity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
		{
			nums: []int{1, 0, 3, 5},
			want: []int{0, 1, 3, 5},
		},
	}

	for _, test := range tests {
		got := sortArrayByParity(test.nums)

		for _, item := range got {
			if !slices.Contains(test.want, item) {
				t.Errorf("sortArrayByParity, result array %v not contain %v", test.want, item)
			}
		}
	}
}

func TestHeightChecker(t *testing.T) {
	t.Parallel()

	tests := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}

	for _, test := range tests {
		if got := heightChecker(test.heights); got != test.want {
			t.Errorf("heightChecker(%v) = %v, want %v", test.heights, got, test.want)
		}
	}
}

func TestThirdMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			nums: []int{2, 2, 3, 3, 3, 4, 4, 4, 4},
			want: 2,
		},
		{
			nums: []int{5, 2, 2},
			want: 5,
		},
		{
			nums: []int{1, 2, 2, 5, 3, 5},
			want: 2,
		},
	}

	for _, test := range tests {
		if got := thirdMax(test.nums); got != test.want {
			t.Errorf("thirdMax(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			nums: []int{1, 1},
			want: []int{2},
		},
		{
			nums: []int{2, 2},
			want: []int{1},
		},
		{
			nums: []int{1, 1, 2, 2},
			want: []int{3, 4},
		},
	}

	for _, test := range tests {
		if got := findDisappearedNumbers(test.nums); !slices.Equal(got, test.want) {
			t.Errorf("findDisappearedNumbers(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
