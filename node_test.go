package red_black_trees

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode(1)
	NotNull(node)
}

func TestNode_IsRed(t *testing.T) {
	n := NewNode(1)
	IsTrue(n.IsRed())
}

func TestNode_Left(t *testing.T) {
	n1 := NewNode(1)
	IsTrue(n1.Left() == nil)
	n2 := NewNode(2)
	n1.SetLeft(n2)
	IsTrue(n1.Left().val == 2)
}

func TestNode_Right(t *testing.T) {
	n1 := NewNode(1)
	IsTrue(n1.Right() == nil)
	n2 := NewNode(2)
	n1.SetRight(n2)
	IsTrue(n1.Right().val == 2)
}

func TestNode_Val(t *testing.T) {
	n1 := NewNode(1)
	IsTrue(n1.Val() == 1)
}
