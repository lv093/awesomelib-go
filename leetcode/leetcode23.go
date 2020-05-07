package main

import "fmt"

func main() {
	node03 := &ListNode{5, nil}
	node02 := &ListNode{4, node03}
	node01 := &ListNode{1, node02}


	node13 := &ListNode{4, nil}
	node12 := &ListNode{3, node13}
	node11 := &ListNode{1, node12}

	node22 := &ListNode{6, nil}
	node21 := &ListNode{2, node22}

	head := mergeKLists([]*ListNode{nil, node01, node11, node21})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	if lists == nil || len(lists) == 0 {
		return head
	}
	if len(lists) == 1 {
		return lists[0]
	}
	head = lists[0]
	idx := 1
	for idx < len(lists) {
		head = mergeTwoNode(head, lists[idx])
		idx++
	}
	return head
}

func mergeTwoNode(node1 *ListNode, node2 *ListNode) (*ListNode) {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	var head, tail *ListNode
	for node1 != nil || node2 != nil {
		var node *ListNode
		if node1 == nil {
			node = &ListNode{node2.Val, nil}
			node2 = node2.Next
		} else if node2 == nil {
			node = &ListNode{node1.Val, nil}
			node1 = node1.Next
		} else {
			if node1.Val <= node2.Val {
				node = &ListNode{node1.Val, nil}
				node1 = node1.Next
			} else {
				node = &ListNode{node2.Val, nil}
				node2 = node2.Next
			}
		}
		if head == nil {
			head = node
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