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

// 				 11b                   11b             11b
//            4r    20r             4r    20r        4b   20b
//          2b   6b         ->    2b  6b       ->   2r 6b
//        1r                    1r 3r             1b 3b
//                                               0r
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

	r.Insert(0)
	AssertIsTrue(r.Find(1).IsBlack())
	AssertIsTrue(r.Find(3).IsBlack())
	AssertIsTrue(r.Find(2).IsRed())
	AssertIsTrue(r.Find(6).IsBlack())
	AssertIsTrue(r.Find(4).IsBlack())
	AssertIsTrue(r.Find(20).IsBlack())
	AssertIsTrue(r.Find(11).IsBlack())
}

//  	   	 11b                  11b
//        4b     20b            6b   20b
//      2r   6b        ->     4b 7r
//    1b 3b 5r 7r		    2r 5r
func Test_LeftRotate(t *testing.T) {
	r := NewRbTree(11)
	r.Root().SetLeft(NewRedNode(4))
	r.Root().SetRight(NewRedNode(20))
	r.Root().Left().SetLeft(NewBlackNode(2))
	r.Root().Left().SetRight(NewBlackNode(6))
	r.Root().Left().Left().SetLeft(NewRedNode(1))
	r.Root().Left().Right().SetLeft(NewRedNode(5))
	r.Root().Left().Right().SetRight(NewRedNode(7))
	r.Insert(3)

	AssertIsTrue(leftRotate(r.Find(4)))
	n6 := r.Find(6)
	AssertIsTrue(IsRoot(n6.Parent()))
	AssertIsTrue(n6.Left().Val() == 4)
	AssertIsTrue(n6.Right().Val() == 7)
	AssertIsTrue(n6.Left().Left().Val() == 2)
	AssertIsTrue(n6.Left().Right().Val() == 5)

	AssertIsTrue(!leftRotate(r.Find(20)))
}

//  	   	 11b                   11b
//        4b     20b             2r   20b
//      2r   6b        ->      1b 4b
//    1b 3b     		     	 3b 6b
func Test_RightRotate(t *testing.T) {
	r := NewRbTree(11)
	r.Root().SetLeft(NewRedNode(4))
	r.Root().SetRight(NewRedNode(20))
	r.Root().Left().SetLeft(NewBlackNode(2))
	r.Root().Left().SetRight(NewBlackNode(6))
	r.Root().Left().Left().SetLeft(NewRedNode(1))
	r.Insert(3)

	AssertIsTrue(rightRotate(r.Find(4)))
	AssertNotTrue(rightRotate(r.Find(20)))
	AssertNotTrue(rightRotate(r.Find(1)))
	AssertIsTrue(r.Find(2).Left().Val() == 1)
	AssertIsTrue(r.Find(2).Right().Val() == 4)
	AssertIsTrue(r.Find(4).Left().Val() == 3)
	AssertIsTrue(r.Find(4).Right().Val() == 6)
}
