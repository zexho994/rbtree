package red_black_trees

import "testing"

func TestNewNode(t *testing.T) {
	node := newRedNode(1)
	AssertNotNull(node)
	AssertNotNull(node.left().parent())
	AssertNotNull(node.right().parent())

	node = newBlackNode(1)
	AssertNotNull(node)
	AssertNotNull(node.left().parent())
	AssertNotNull(node.right().parent())
}

func TestNode_IsBlack(t *testing.T) {
	n := newBlackNode(1)
	AssertIsTrue(n.isBlack())
}

func TestNode_IsRed(t *testing.T) {
	n := newRedNode(1)
	AssertIsTrue(n.isRed())
}

func TestNode_Left(t *testing.T) {
	n1 := newBlackNode(1)
	n2 := newBlackNode(2)
	n1.setLeft(n2)
	AssertIsTrue(n1.left().val() == 2)
	AssertIsTrue(n1.left().parent() == n1)
	AssertIsTrue(!n1.left().isLeaf())
	AssertIsTrue(n1.left().left().isLeaf())
	AssertIsTrue(!n1.left().isRed())
}

func TestNode_Right(t *testing.T) {
	n1 := newBlackNode(1)
	n2 := newBlackNode(2)
	n1.setRight(n2)

	AssertIsTrue(n1.right().val() == 2)
	AssertIsTrue(n1.right().parent() == n1)
	AssertIsTrue(!n1.right().isLeaf())
	AssertIsTrue(n1.right().right().isLeaf())
	AssertIsTrue(!n1.right().isRed())
}

func TestNode_Val(t *testing.T) {
	n1 := newRedNode(1)
	AssertIsTrue(n1.val() == 1)
}

func TestNode_Parent(t *testing.T) {
	n1 := newBlackNode(2)
	n2 := newBlackNode(1)
	n3 := newBlackNode(3)
	n1.setLeft(n2)
	n1.setRight(n3)

	AssertIsTrue(n1.left().parent() == n1)
	AssertIsTrue(n1.right().parent() == n1)
}
