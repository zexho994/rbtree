package red_black_trees

import "testing"

func TestNewRbTree(t *testing.T) {
	t1 := NewRbTree(1)
	IsTrue(t1.root.IsRed())
}

func TestRbTree_Find(t *testing.T) {
	t1 := NewRbTree(2)
	t1.Root().SetLeft(NewNode(1))
	t1.Root().SetRight(NewNode(3))
	NotNull(t1.Find(1))
	NotNull(t1.Find(3))
}
