package base_algo

import "fmt"

/**
1.链表：头结点、尾结点。
2.结点：前一个结点、后一个结点、结点值
3.结点类型
*/


type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val interface{}
	Previous *Node
	Next *Node
}

type NodeValue struct {
	Value int
}

func (this *LinkedList) PrintList() {
	curr := this.Head

	for curr != nil {
		//print current node
		curr.PrintNode()
		curr = curr.Next
	}
}

func (this *Node) PrintNode() {
	fmt.Print(this.Val)
}

func (this *LinkedList) InsertHead(node *Node) {
	node.Next = this.Head
	this.Head = node
}

func (this *LinkedList)Insert(node *Node) {

}

func (this *LinkedList) Exist(value interface{}) {

}

func (this *LinkedList) Remove(node *Node) {

}


//反转单链表,会破坏原链表
func ReverseList(head *Node) *Node {
	cur := head
	fmt.Println("head===", head)
	fmt.Println("head===curr==", cur)
	var pre *Node
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
//
//func MergeOrderList(node1, node2 *Node) *Node {
//	curr1, curr2 := node1, node2
//
//	var head, next Node
//	for curr1 != nil || curr2 != nil {
//		if curr1 == nil {
//
//		} else if curr2 == nil {
//
//		} else {
//
//		}
//	}
//}

/**
单链表反转
 */
func (this *LinkedList) Reverse() *LinkedList {
	newLink := new(LinkedList)

	if this.Head == nil {
		return newLink
	}
	next := this.Head.Next
	if next == nil {
		return this
	}
	newLink.Tail = this.Head

	oldCurr := this.Head
	var oldPrev *Node = nil
	for oldCurr != nil {
		if oldPrev == nil {
			head := new(Node)
			head.Val = oldCurr.Val
			head.Next = nil
			newLink.Head = head
		} else {
			node := new(Node)
			node.Val = oldCurr.Val
			newLink.InsertHead(node)
		}

		oldPrev = oldCurr
		oldCurr = oldCurr.Next
	}

	return newLink
}