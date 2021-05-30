package red_black_trees

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode(1)
	AssertNotNull(node)
}

func TestNode_IsRed(t *testing.T) {
	n := NewNode(1)
	AssertIsTrue(n.IsRed())
}

func TestNode_Left(t *testing.T) {
	n1 := NewNode(1)
	AssertIsTrue(n1.Left() == nil)
	n2 := NewNode(2)
	n1.SetLeft(n2)
	AssertIsTrue(n1.Left().Val() == 2)
	AssertIsTrue(n1.Left().Parent() == n1)
}

func TestNode_Right(t *testing.T) {
	n1 := NewNode(1)
	AssertIsTrue(n1.Right() == nil)
	n2 := NewNode(2)
	n1.SetRight(n2)

	AssertIsTrue(n1.Right().Val() == 2)
	AssertIsTrue(n1.Right().Parent() == n1)
}

func TestNode_Val(t *testing.T) {
	n1 := NewNode(1)
	AssertIsTrue(n1.Val() == 1)
}

func TestNode_Parent(t *testing.T) {
	n1 := NewNode(2)
	AssertIsTrue(n1.Right() == nil)
	n2 := NewNode(1)
	n3 := NewNode(3)
	n1.SetLeft(n2)
	n1.SetRight(n3)

	AssertIsTrue(n1.Left().Parent() == n1)
	AssertIsTrue(n1.Right().Parent() == n1)
}
