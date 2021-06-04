package rbtree

import "testing"

func TestNewRbTree(t *testing.T) {
	t1 := NewRbTree(1)
	AssertIsTrue(!t1.r.isRed())
}

func TestRbTree_Find(t *testing.T) {
	t1 := NewRbTree(2)
	t1.root().setLeft(newRedNode(1))
	t1.root().setRight(newRedNode(3))
	AssertNotNull(t1.Find(1))
	AssertNotNull(t1.Find(3))
	AssertIsTrue(!t1.Find(3).isLeafNode())
	AssertIsTrue(t1.Find(4).isLeafNode())
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

	// check v
	AssertIsTrue(rbt.Find(11).left().val() == 2)
	AssertIsTrue(rbt.Find(11).right().val() == 25)
	AssertIsTrue(rbt.Find(2).left().val() == 1)
	AssertIsTrue(rbt.Find(2).right().val() == 4)
	AssertIsTrue(rbt.Find(25).left().val() == 20)
	AssertIsTrue(rbt.Find(25).right().val() == 50)

	// check c
	AssertIsTrue(rbt.Find(11).left().isBlack())
	AssertIsTrue(rbt.Find(11).right().isBlack())
	AssertIsTrue(rbt.Find(2).left().isRed())
	AssertIsTrue(rbt.Find(2).right().isRed())
	AssertIsTrue(rbt.Find(25).left().isRed())
	AssertIsTrue(rbt.Find(25).right().isRed())
}

func TestRbTree_Insert2(t *testing.T) {
	rbt := NewRbTree(10)
	rbt.Insert(5)
	rbt.Insert(1)
	rbt.Insert(0)

	AssertIsTrue(isRoot(rbt.Find(5)))
	AssertIsTrue(rbt.Find(5).left().val() == 1)
	AssertIsTrue(rbt.Find(5).right().val() == 10)
}

//  	   	 11b                  11b
//        4b     20b            6b   20b
//      2r   6b        ->     4b 7r
//    1b 3b 5r 7r		    2r 5r
func Test_LeftRotate(t *testing.T) {
	r := NewRbTree(11)
	r.root().setLeft(newRedNode(4))
	r.root().setRight(newRedNode(20))
	r.root().left().setLeft(newBlackNode(2))
	r.root().left().setRight(newBlackNode(6))
	r.root().left().left().setLeft(newRedNode(1))
	r.root().left().right().setLeft(newRedNode(5))
	r.root().left().right().setRight(newRedNode(7))
	r.Insert(3)

	AssertIsTrue(r.leftRotate(r.Find(4)))
	n6 := r.Find(6)
	AssertIsTrue(isRoot(n6.parent()))
	AssertIsTrue(n6.left().val() == 4)
	AssertIsTrue(n6.right().val() == 7)
	AssertIsTrue(n6.left().left().val() == 2)
	AssertIsTrue(n6.left().right().val() == 5)

	AssertIsTrue(!r.leftRotate(r.Find(20)))
}

//  	   	 11b                   11b
//        4b     20b             2r   20b
//      2r   6b        ->      1b 4b
//    1b 3b     		     	 3b 6b
func Test_RightRotate(t *testing.T) {
	r := NewRbTree(11)
	r.root().setLeft(newRedNode(4))
	r.root().setRight(newRedNode(20))
	r.root().left().setLeft(newBlackNode(2))
	r.root().left().setRight(newBlackNode(6))
	r.root().left().left().setLeft(newRedNode(1))
	r.Insert(3)

	AssertIsTrue(r.rightRotate(r.Find(4)))
	AssertNotTrue(r.rightRotate(r.Find(20)))
	AssertNotTrue(r.rightRotate(r.Find(1)))
	AssertIsTrue(r.Find(2).left().val() == 1)
	AssertIsTrue(r.Find(2).right().val() == 4)
	AssertIsTrue(r.Find(4).left().val() == 3)
	AssertIsTrue(r.Find(4).right().val() == 6)
}

func TestIsRbTree(t *testing.T) {
	rbt := NewRbTree(11)
	isRbtree(rbt.root(), 0)
	rbt.Insert(4)
	isRbtree(rbt.root(), 0)
	rbt.Insert(20)
	isRbtree(rbt.root(), 0)
	rbt.Insert(2)
	isRbtree(rbt.root(), 0)
	rbt.Insert(1)
	isRbtree(rbt.root(), 0)
	rbt.Insert(50)
	isRbtree(rbt.root(), 0)
	rbt.Insert(25)
	isRbtree(rbt.root(), 0)
	rbt.Insert(132)
	isRbtree(rbt.root(), 0)
	rbt.Insert(42)
	isRbtree(rbt.root(), 0)
	rbt.Insert(58)
	isRbtree(rbt.root(), 0)
	rbt.Insert(53)
	isRbtree(rbt.root(), 0)
	rbt.Insert(70)
	isRbtree(rbt.root(), 0)
	rbt.Insert(-3)
	isRbtree(rbt.root(), 0)
	rbt.Insert(6)
	isRbtree(rbt.root(), 0)
	rbt.Insert(20)
	isRbtree(rbt.root(), 0)
	rbt.Insert(-10)
	isRbtree(rbt.root(), 0)
	rbt.Insert(100)
	isRbtree(rbt.root(), 0)
}

func isRbtree(n *node, s int) int {
	if n.isLeafNode() {
		return s
	}
	if n.isRed() {
		if n.parent().isRed() {
			panic("not is RbTree")
		}
	} else {
		s++
	}
	c1 := isRbtree(n.l, s)
	c2 := isRbtree(n.r, s)
	if c1 != c2 {
		panic("not is RbTree")
	}
	return c1

}

func TestRbTree_remove(t1 *testing.T) {
	rbt := NewRbTree(11)
	rbt.Insert(4)
	rbt.Insert(20)
	rbt.Insert(2)
	rbt.Insert(1)
	rbt.Insert(50)
	rbt.Insert(25)
	rbt.Insert(132)

	rbt.Remove(25)
	isRbtree(rbt.root(), 0)
	rbt.Remove(50)
	isRbtree(rbt.root(), 0)
	rbt.Remove(1)
	isRbtree(rbt.root(), 0)
	rbt.Remove(20)
	isRbtree(rbt.root(), 0)
	rbt.Insert(40)
	isRbtree(rbt.root(), 0)
	rbt.Remove(2)
	isRbtree(rbt.root(), 0)
	rbt.Insert(46)
	isRbtree(rbt.root(), 0)
	rbt.Remove(11)
	isRbtree(rbt.root(), 0)
	rbt.Insert(59)
	isRbtree(rbt.root(), 0)
}
