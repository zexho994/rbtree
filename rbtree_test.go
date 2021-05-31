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

//            11b
//        2r       25b
//      1b   4b  20r  50r
func TestRbTree_Insert(t *testing.T) {
	rbt := NewRbTree(11)
	rbt.Insert(4)
	rbt.Insert(20)
	rbt.Insert(2)
	rbt.Insert(1)
	rbt.Insert(50)
	rbt.Insert(25)

	// check val
	AssertIsTrue(rbt.Find(11).Left().Val() == 2)
	AssertIsTrue(rbt.Find(11).Right().Val() == 25)
	AssertIsTrue(rbt.Find(2).Left().Val() == 1)
	AssertIsTrue(rbt.Find(2).Right().Val() == 4)
	AssertIsTrue(rbt.Find(25).Left().Val() == 20)
	AssertIsTrue(rbt.Find(25).Right().Val() == 50)

	// check color
	AssertIsTrue(rbt.Find(11).Left().IsBlack())
	AssertIsTrue(rbt.Find(11).Right().IsBlack())
	AssertIsTrue(rbt.Find(2).Left().IsRed())
	AssertIsTrue(rbt.Find(2).Right().IsRed())
	AssertIsTrue(rbt.Find(25).Left().IsRed())
	AssertIsTrue(rbt.Find(25).Right().IsRed())
}

func TestRbTree_Insert2(t *testing.T) {
	rbt := NewRbTree(10)
	rbt.Insert(5)
	rbt.Insert(1)
	rbt.Insert(0)

	AssertIsTrue(IsRoot(rbt.Find(5)))
	AssertIsTrue(rbt.Find(5).Left().Val() == 1)
	AssertIsTrue(rbt.Find(5).Right().Val() == 10)
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

	AssertIsTrue(r.leftRotate(r.Find(4)))
	n6 := r.Find(6)
	AssertIsTrue(IsRoot(n6.Parent()))
	AssertIsTrue(n6.Left().Val() == 4)
	AssertIsTrue(n6.Right().Val() == 7)
	AssertIsTrue(n6.Left().Left().Val() == 2)
	AssertIsTrue(n6.Left().Right().Val() == 5)

	AssertIsTrue(!r.leftRotate(r.Find(20)))
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

	AssertIsTrue(r.rightRotate(r.Find(4)))
	AssertNotTrue(r.rightRotate(r.Find(20)))
	AssertNotTrue(r.rightRotate(r.Find(1)))
	AssertIsTrue(r.Find(2).Left().Val() == 1)
	AssertIsTrue(r.Find(2).Right().Val() == 4)
	AssertIsTrue(r.Find(4).Left().Val() == 3)
	AssertIsTrue(r.Find(4).Right().Val() == 6)
}

func TestIsRbTree(t *testing.T) {
	rbt := NewRbTree(11)
	rbt.Insert(4)
	rbt.Insert(20)
	rbt.Insert(2)
	rbt.Insert(1)
	rbt.Insert(50)
	//rbt.Insert(25)
	rbt.Insert(132)
	rbt.Insert(42)
	rbt.Insert(58)
	rbt.Insert(53)
	rbt.Insert(70)
	rbt.Insert(16)

	var isRbtree func(n *node, s int) int
	isRbtree = func(n *node, s int) int {
		if n.IsLeaf() {
			return s
		}
		if n.IsRed() {
			if n.Parent().IsRed() {
				panic("not is RbTree")
			}
		} else {
			s++
		}
		c1 := isRbtree(n.left, s)
		c2 := isRbtree(n.right, s)
		if c1 != c2 {
			panic("not is RbTree")
		}
		return c1
	}
	isRbtree(rbt.Root(), 0)
}
