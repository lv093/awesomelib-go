package main

func main() {
	head := &ListNode{1, &ListNode{}}
	node2 := &ListNode{2, &ListNode{}}
	node3 := &ListNode{3, &ListNode{}}
	node4 := &ListNode{4, &ListNode{}}
	node5 := &ListNode{5, &ListNode{}}
	node6 := &ListNode{6, nil}

	node5.Next = node6
	node4.Next = node5
	node3.Next = node4
	node2.Next = node3
	head.Next = node2

	head1 := mergeTwoLists(head, nil)
	for head1 != nil {
		head1 = head1.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	for l1 != nil || l2 != nil {

		var node *ListNode
		if l1 == nil {
			node = l2
			l2 = l2.Next
		} else if l2 == nil {
			node = l1
			l1 = l1.Next
		} else if l1.Val >= l2.Val {
			node = l2
			l2 = l2.Next
		} else {
			node = l1
			l1 = l1.Next
		}
		if head == nil {
			head = node
			head.Next = tail
		} else if tail == nil {
			tail = node
			head.Next = tail
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}