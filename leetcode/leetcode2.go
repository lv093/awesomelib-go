package main

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{9, &ListNode{9, nil}}
	res := addTwoNumbers(l1, l2)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1, curr2 := l1, l2
	step := 0
	var root, curr *ListNode
	for curr1 != nil || curr2 != nil {
		val1, val2 := 0, 0
		if curr1 != nil{
			val1 = curr1.Val
		}
		if curr2 != nil {
			val2 = curr2.Val
		}
		node := &ListNode{Val:(val1 + val2 + step) % 10, Next:nil}
		step = (val1 + val2 + step) / 10
		if root == nil {
			root = node
			curr = node
		} else {
			curr.Next = node
			curr = curr.Next
		}

		if curr1 != nil {
			curr1 = curr1.Next
		}
		if curr2 != nil {
			curr2 = curr2.Next
		}
	}
	for step != 0 {
		curr.Next = &ListNode{Val: step % 10, Next:nil}
		step = step / 10
	}
	return root
}
