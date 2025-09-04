package linkedlist

import (
	"testing"
)

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

func TestLinkedList1(t *testing.T) {
	t.Parallel()

	ll := Constructor()
	ll.AddAtIndex(2, 1)
	ll.AddAtIndex(3, 4)
	ll.AddAtTail(1)

	if got := ll.Get(0); got != 1 {
		t.Errorf("TestLinkedList1 Get - want %v, got %v", -1, got)
	}

	ll.DeleteAtIndex(0)

	if got := ll.Get(0); got != -1 {
		t.Errorf("TestLinkedList1 Get - want %v, got %v", 1, got)
	}
}

func TestLinkedList2(t *testing.T) {
	t.Parallel()

	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtTail(3)
	ll.AddAtIndex(1, 2)

	if got := ll.Get(1); got != 2 {
		t.Errorf("TestLinkedList2 Get - want %v, got %v", 2, got)
	}

	ll.DeleteAtIndex(1)

	if got := ll.Get(1); got != 3 {
		t.Errorf("TestLinkedList2 Get - want %v, got %v", 3, got)
	}
}
