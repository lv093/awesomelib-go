package main

const (
	RED = true
	BLACK = false
)

type RBNode struct {
	color	bool
	Data 	interface{}
	Parent 	*RBNode
	Left 	*RBNode
	Right 	*RBNode
}

func (this *RBNode) getParentNode() *RBNode {
	//TODO
	return this
}



type RBTree struct {
	Root 	*RBNode
}

func (this *RBTree) Insert(target *RBNode) bool {

}

func (this *RBTree) rotate(x *RBNode, isLeftRotate bool) {
	if isLeftRotate {
		right := x.Right
		if right == nil {
			return
		}
		parent := x.Parent

		x.Right = right.Left
		right.Left.Parent = x
		right.Parent = parent

		if parent.Left == x {
			parent.Left = right
		} else if parent.Right == x {
			parent.Right = right
		} else {
			this.Root = right
		}

		right.Left = x
		x.Parent = right
	} else {
		left := x.Left
		if left == nil {
			return
		}
		parent := x.Parent

		x.Left = left.Right
		left.Right.Parent = x
		left.Parent = parent

		if parent.Left == x {
			parent.Left = left
		} else if parent.Right == x {
			parent.Right = left
		} else {
			this.Root = left
		}

		left.Right = x
		x.Parent = left
	}
}