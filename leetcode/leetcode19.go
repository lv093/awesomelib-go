package main

import "fmt"

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

	node := removeNthFromEnd(head, 6)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	delHead := head
	var lastDelHead *ListNode
	delLen := 1
	for node != nil {
		if delLen <= n {
			delLen++
		} else {
			lastDelHead = delHead
			delHead = delHead.Next
		}
		node = node.Next
	}
	if lastDelHead == nil {
		head = head.Next
	} else {
		fmt.Println(delHead)
		lastDelHead.Next = delHead.Next

	}
	return head
}