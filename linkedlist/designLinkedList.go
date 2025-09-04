package linkedlist

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Length int
	Head   *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.Length-1 < index {
		return -1
	}

	countStep := 0
	currentNode := this.Head

	for countStep < index {
		currentNode = currentNode.Next
		countStep++
	}

	return currentNode.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val, Next: nil}

	this.Length++

	if this.Head == nil {
		this.Head = newNode
		return
	}

	newNode.Next = this.Head
	this.Head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{Val: val, Next: nil}

	this.Length++

	if this.Head == nil {
		this.Head = newNode
		return
	}

	currentNode := this.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.Length == index {
		this.AddAtTail(val)
		return
	}

	if this.Length-1 < index {
		return
	}

	newNode := &Node{Val: val, Next: nil}
	this.Length++

	itemIdx := 0
	currentNode := this.Head

	for itemIdx+1 < index {
		currentNode = currentNode.Next
		itemIdx++
	}

	newNode.Next = currentNode.Next
	currentNode.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	switch {
	case this.Head == nil:
		return
	case this.Length-1 < index:
		return
	case index == 0:
		curr := this.Head
		this.Head = curr.Next
		this.Length--
		return
	case this.Length-1 == index:
		curr := this.Head

		for curr.Next.Next != nil {
			curr = curr.Next
		}

		curr.Next = nil
		this.Length--
		return
	}

	var prev *Node
	curr := this.Head
	count := 0

	for count < index {
		prev = curr
		curr = curr.Next
		count++
	}

	prev.Next = curr.Next
	this.Length--
}
