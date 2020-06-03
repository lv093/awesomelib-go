package main

import "fmt"

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
 */
func main() {
	n4 := &ListNode{4, nil}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	ne := swapPairs(n1)
	fmt.Println(ne, ne.Next, ne.Next.Next, ne.Next.Next.Next)
}
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
通过在原链表上进行翻转
在内存消耗和时间上都是最优。
 */
func swapPairs(head *ListNode) *ListNode {
	last := head
	if head == nil {
		return head
	}
	curr := last.Next

	if curr == nil {
		return head
	}
	next := curr.Next
	//swap
	curr.Next = last
	last.Next = next
	head = curr

	prev := last
	for next != nil {
		first := next
		second := next.Next

		if second == nil {
			break
		}
		prev.Next = second
		first.Next = second.Next
		next = second.Next
		second.Next = first

		prev = first
	}
	return head
}
 /**
 方法：通过复制一条新链
  */
func swapPairs1(head *ListNode) *ListNode {
	var newHead, newTail *ListNode
	n := head

	for n != nil {
		first := ListNode{n.Val, nil}
		if newHead == nil {
			newHead = &first
		}
		if n.Next == nil {
			if newTail == nil {
				break
			} else {
				newTail.Next = &first
				break
			}
		}
		second := ListNode{n.Next.Val, nil}
		second.Next = &first
		if newTail == nil {
			newHead = &second
			newTail = &first
		} else {
			newTail.Next = &second
			newTail = &first
		}
		n = n.Next.Next
	}
	return newHead
}

