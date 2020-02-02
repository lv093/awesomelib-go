package base_algo

import "testing"

func TestLinedList(t *testing.T) {
	linkedList := new(LinkedList)

	node1 := new(Node)
	nodeValue := NodeValue{1}
	node1.Val = nodeValue

	node2 := new(Node)
	nodeValue.Value = 2
	node2.Val = nodeValue

	node3 := new(Node)
	nodeValue.Value = 3
	node3.Val = nodeValue

	node4 := new(Node)
	nodeValue.Value = 4
	node4.Val = nodeValue

	linkedList.Head = node1
	node1.Next = node2

	node2.Next = node3
	node3.Next = node4

	linkedList.Tail = node4
	linkedList.PrintList()

	ReverseList(linkedList.Head).PrintNode()

	linkedList.PrintList()
}