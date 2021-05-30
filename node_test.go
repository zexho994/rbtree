package red_black_trees

import "testing"

func TestNewNode(t *testing.T) {
	node := NewRedNode(1)
	AssertNotNull(node)
	AssertNotNull(node.Left().Parent())
	AssertNotNull(node.Right().Parent())

	node = NewBlackNode(1)
	AssertNotNull(node)
	AssertNotNull(node.Left().Parent())
	AssertNotNull(node.Right().Parent())
}

func TestNode_IsBlack(t *testing.T) {
	n := NewBlackNode(1)
	AssertIsTrue(n.IsBlack())
}

func TestNode_IsRed(t *testing.T) {
	n := NewRedNode(1)
	AssertIsTrue(n.IsRed())
}

func TestNode_Left(t *testing.T) {
	n1 := NewBlackNode(1)
	n2 := NewBlackNode(2)
	n1.SetLeft(n2)
	AssertIsTrue(n1.Left().Val() == 2)
	AssertIsTrue(n1.Left().Parent() == n1)
	AssertIsTrue(!n1.Left().IsLeaf())
	AssertIsTrue(n1.Left().Left().IsLeaf())
	AssertIsTrue(!n1.Left().IsRed())
}

func TestNode_Right(t *testing.T) {
	n1 := NewBlackNode(1)
	n2 := NewBlackNode(2)
	n1.SetRight(n2)

	AssertIsTrue(n1.Right().Val() == 2)
	AssertIsTrue(n1.Right().Parent() == n1)
	AssertIsTrue(!n1.Right().IsLeaf())
	AssertIsTrue(n1.Right().Right().IsLeaf())
	AssertIsTrue(!n1.Right().IsRed())
}

func TestNode_Val(t *testing.T) {
	n1 := NewRedNode(1)
	AssertIsTrue(n1.Val() == 1)
}

func TestNode_Parent(t *testing.T) {
	n1 := NewBlackNode(2)
	n2 := NewBlackNode(1)
	n3 := NewBlackNode(3)
	n1.SetLeft(n2)
	n1.SetRight(n3)

	AssertIsTrue(n1.Left().Parent() == n1)
	AssertIsTrue(n1.Right().Parent() == n1)
}
