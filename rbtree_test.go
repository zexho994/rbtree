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

// 				11                  11
//            4    20             4    20
//          2   6         ->    2  6
//        1                    1 3
func TestRbTree_Insert(t *testing.T) {
	r := NewRbTree(11)
	r.Root().SetLeft(NewRedNode(4))
	r.Root().SetRight(NewRedNode(20))
	r.Root().Left().SetLeft(NewBlackNode(2))
	r.Root().Left().SetRight(NewBlackNode(6))
	r.Root().Left().Left().SetLeft(NewRedNode(1))
	AssertIsTrue(r.Find(2).IsNonLeaf())
	AssertIsTrue(r.Find(3).IsLeaf())

	r.Insert(3)
	f := r.Find(3)
	AssertIsTrue(f.IsNonLeaf())
	AssertIsTrue(f.Grandfather().Val() == 4)
	AssertIsTrue(f.Uncle().Val() == 6)
	AssertIsTrue(r.Find(4).Uncle() == nil)
	AssertIsTrue(r.Find(11).Uncle() == nil)
}
