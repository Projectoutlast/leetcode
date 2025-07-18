package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var nodes []*ListNode

	for curr := head; curr != nil; curr = curr.Next {
		nodes = append(nodes, curr)
	}

	return nodes[len(nodes)/2]
}

func fastMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
