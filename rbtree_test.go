package red_black_trees

import "testing"

func TestNewRbTree(t *testing.T) {
	t1 := NewRbTree(1)
	AssertIsTrue(!t1.root.IsRed())
}

func TestRbTree_Find(t *testing.T) {
	t1 := NewRbTree(2)
	t1.Root().SetLeft(NewRedNode(1))
	t1.Root().SetRight(NewRedNode(3))
	AssertNotNull(t1.Find(1))
	AssertNotNull(t1.Find(3))
	AssertIsTrue(!t1.Find(3).IsLeaf())
	AssertIsTrue(t1.Find(4).IsLeaf())
}
