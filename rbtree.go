package red_black_trees

// rb.node is do not allow duplication
type rbTree struct {
	r *node // root
}

func newRbTree(v V) *rbTree {
	return &rbTree{
		r: newBlackNode(v),
	}
}

func (t *rbTree) root() *node {
	return t.r
}

// find 值为v的节点
// 如果存在，返回该节点，该节点肯定不是叶子节点
// 否则，返回的节点 isLeaf() == true
func (t *rbTree) find(v V) *node {
	m := t.root()
	for m.isNonLeaf() {
		if v > m.val() {
			m = m.right()
		} else if v < m.val() {
			m = m.left()
		} else {
			return m
		}
	}

	return m
}

// insert a node into rbTree
func (t *rbTree) insert(v V) {
	n := t.find(v)
	if n.isNonLeaf() {
		return
	}
	// 保存新的节点到对应的位置
	if parent := n.parent(); v < parent.val() {
		parent.setLeft(newRedNode(v))
		n = parent.left()
	} else {
		parent.setRight(newRedNode(v))
		n = parent.right()
	}

CASE1:
	if n.parent().isBlack() {
		return
	}
	if n.uncle() == nil || n.uncle().isBlack() {
		goto CASE2
	}

	n.parent().turnBlack()
	n.uncle().turnBlack()
	if isRoot(n.grandfather()) {
		return
	}
	n.grandfather().turnRed()
	n = n.grandfather()
	goto CASE1

CASE2: // 当 node父节点的关系 != node父节点与祖父节点的关系 时成立
	if (n.isLeft() && n.parent().isLeft()) || (n.isRight() && n.parent().isRight()) {
		goto CASE3
	}
	if n.isLeft() {
		n = n.p
		t.rightRotate(n)
	} else {
		n = n.p
		t.leftRotate(n)
	}

CASE3: // 当 node父节点的关系 == node父节点与祖父节点的关系 时成立
	if n.isLeft() {
		t.rightRotate(n.grandfather())
	} else {
		t.leftRotate(n.grandfather())
	}
	n.parent().turnBlack()
	n.brother().turnRed()
}

// remove a node from rbTree
func (t *rbTree) remove(v V) {
	n := t.find(v)
	if n.isLeaf() {
		return
	}
	if isRoot(n) {
		t.r = nil
		return
	}
	if n.sonCount() == 0 {
		if n.isLeft() {
			n.parent().setLeft(newNilNode())
		} else {
			n.parent().setRight(newNilNode())
		}
		return
	}
	var son *node
	if n.sonCount() == 1 {
		if n.left().isNonLeaf() {
			son = n.left()
		} else {
			son = n.right()
		}
		son.turnBlack()
		if n.isLeft() {
			n.parent().setLeft(n.right())
		} else {
			n.parent().setRight(n.right())
		}
		n.setParent(nil)
		return
	}

	// n.sonCount() == 2
	son = n.right()
	for son.left().isNonLeaf() {
		son = son.left()
	}

}

func (t *rbTree) delNode(n *node) {

}

func isRoot(n *node) bool {
	if n == nil {
		return false
	}
	return n.p == nil
}

// 1. x父指针的子指针指向c
// 2. x的右子指针指向c的左子节点
// 3. c的左子指针指向x
func (t *rbTree) leftRotate(x *node) bool {
	if x == nil || x.isLeaf() || x.right().isLeaf() {
		return false
	}
	c := x.right()
	if t.root() == x {
		t.setRoot(c)
		x.setRight(c.left())
		c.setLeft(x)
		return true
	}

	if x.isLeft() {
		x.parent().setLeft(c)
	} else {
		x.parent().setRight(c)
	}
	x.setRight(c.left())
	c.setLeft(x)
	return true
}

// 1.x父节点的子指针指向x左节点B
// 2.x左节点指向B右节点
// 3.B右节点指向x
func (t *rbTree) rightRotate(x *node) bool {
	if x == nil || x.isLeaf() || x.left().isLeaf() {
		return false
	}
	b := x.left()
	if t.root() == x {
		t.setRoot(b)
		x.setLeft(b.right())
		b.setRight(x)
		return true
	}
	if x.isLeft() {
		x.parent().setLeft(b)
	} else {
		x.parent().setRight(b)
	}
	x.setLeft(b.right())
	b.setRight(x)
	return true
}

func (t *rbTree) setRoot(n *node) {
	if n == nil {
		return
	}
	n.p = nil
	t.r = n
}
