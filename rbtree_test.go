package red_black_trees

import "testing"

func TestNewRbTree(t *testing.T) {
	t1 := newRbTree(1)
	AssertIsTrue(!t1.r.isRed())
}

func TestRbTree_Find(t *testing.T) {
	t1 := newRbTree(2)
	t1.root().setLeft(newRedNode(1))
	t1.root().setRight(newRedNode(3))
	AssertNotNull(t1.find(1))
	AssertNotNull(t1.find(3))
	AssertIsTrue(!t1.find(3).isLeafNode())
	AssertIsTrue(t1.find(4).isLeafNode())
}

//            11b
//        2r       25b
//      1b   4b  20r  50r
func TestRbTree_Insert(t *testing.T) {
	rbt := newRbTree(11)
	rbt.insert(4)
	rbt.insert(20)
	rbt.insert(2)
	rbt.insert(1)
	rbt.insert(50)
	rbt.insert(25)

	// check v
	AssertIsTrue(rbt.find(11).left().val() == 2)
	AssertIsTrue(rbt.find(11).right().val() == 25)
	AssertIsTrue(rbt.find(2).left().val() == 1)
	AssertIsTrue(rbt.find(2).right().val() == 4)
	AssertIsTrue(rbt.find(25).left().val() == 20)
	AssertIsTrue(rbt.find(25).right().val() == 50)

	// check c
	AssertIsTrue(rbt.find(11).left().isBlack())
	AssertIsTrue(rbt.find(11).right().isBlack())
	AssertIsTrue(rbt.find(2).left().isRed())
	AssertIsTrue(rbt.find(2).right().isRed())
	AssertIsTrue(rbt.find(25).left().isRed())
	AssertIsTrue(rbt.find(25).right().isRed())
}

func TestRbTree_Insert2(t *testing.T) {
	rbt := newRbTree(10)
	rbt.insert(5)
	rbt.insert(1)
	rbt.insert(0)

	AssertIsTrue(isRoot(rbt.find(5)))
	AssertIsTrue(rbt.find(5).left().val() == 1)
	AssertIsTrue(rbt.find(5).right().val() == 10)
}

//  	   	 11b                  11b
//        4b     20b            6b   20b
//      2r   6b        ->     4b 7r
//    1b 3b 5r 7r		    2r 5r
func Test_LeftRotate(t *testing.T) {
	r := newRbTree(11)
	r.root().setLeft(newRedNode(4))
	r.root().setRight(newRedNode(20))
	r.root().left().setLeft(newBlackNode(2))
	r.root().left().setRight(newBlackNode(6))
	r.root().left().left().setLeft(newRedNode(1))
	r.root().left().right().setLeft(newRedNode(5))
	r.root().left().right().setRight(newRedNode(7))
	r.insert(3)

	AssertIsTrue(r.leftRotate(r.find(4)))
	n6 := r.find(6)
	AssertIsTrue(isRoot(n6.parent()))
	AssertIsTrue(n6.left().val() == 4)
	AssertIsTrue(n6.right().val() == 7)
	AssertIsTrue(n6.left().left().val() == 2)
	AssertIsTrue(n6.left().right().val() == 5)

	AssertIsTrue(!r.leftRotate(r.find(20)))
}

//  	   	 11b                   11b
//        4b     20b             2r   20b
//      2r   6b        ->      1b 4b
//    1b 3b     		     	 3b 6b
func Test_RightRotate(t *testing.T) {
	r := newRbTree(11)
	r.root().setLeft(newRedNode(4))
	r.root().setRight(newRedNode(20))
	r.root().left().setLeft(newBlackNode(2))
	r.root().left().setRight(newBlackNode(6))
	r.root().left().left().setLeft(newRedNode(1))
	r.insert(3)

	AssertIsTrue(r.rightRotate(r.find(4)))
	AssertNotTrue(r.rightRotate(r.find(20)))
	AssertNotTrue(r.rightRotate(r.find(1)))
	AssertIsTrue(r.find(2).left().val() == 1)
	AssertIsTrue(r.find(2).right().val() == 4)
	AssertIsTrue(r.find(4).left().val() == 3)
	AssertIsTrue(r.find(4).right().val() == 6)
}

func TestIsRbTree(t *testing.T) {
	rbt := newRbTree(11)
	isRbtree(rbt.root(), 0)
	rbt.insert(4)
	isRbtree(rbt.root(), 0)
	rbt.insert(20)
	isRbtree(rbt.root(), 0)
	rbt.insert(2)
	isRbtree(rbt.root(), 0)
	rbt.insert(1)
	isRbtree(rbt.root(), 0)
	rbt.insert(50)
	isRbtree(rbt.root(), 0)
	rbt.insert(25)
	isRbtree(rbt.root(), 0)
	rbt.insert(132)
	isRbtree(rbt.root(), 0)
	rbt.insert(42)
	isRbtree(rbt.root(), 0)
	rbt.insert(58)
	isRbtree(rbt.root(), 0)
	rbt.insert(53)
	isRbtree(rbt.root(), 0)
	rbt.insert(70)
	isRbtree(rbt.root(), 0)
	rbt.insert(-3)
	isRbtree(rbt.root(), 0)
	rbt.insert(6)
	isRbtree(rbt.root(), 0)
	rbt.insert(20)
	isRbtree(rbt.root(), 0)
	rbt.insert(-10)
	isRbtree(rbt.root(), 0)
	rbt.insert(100)
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
	rbt := newRbTree(11)
	rbt.insert(4)
	rbt.insert(20)
	rbt.insert(2)
	rbt.insert(1)
	rbt.insert(50)
	rbt.insert(25)
	rbt.insert(132)

	rbt.remove(25)
	isRbtree(rbt.root(), 0)
}
