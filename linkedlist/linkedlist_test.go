package linkedlist

import "testing"

func TestMiddleNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("MiddleNode", func(t *testing.T) {
			if got := middleNode(tt.head); got.Val != tt.want.Val {
				t.Errorf("middleNode() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}

func TestFastMiddleNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("MiddleNode", func(t *testing.T) {
			if got := fastMiddleNode(tt.head); got.Val != tt.want.Val {
				t.Errorf("FastMiddleNode() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
