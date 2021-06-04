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
// 否则，返回的节点 isLeafNode() == true
func (t *rbTree) find(v V) *node {
	m := t.root()
	for m.isNonLeafNode() {
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
	if n.isNonLeafNode() {
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
	f := t.find(v)
	if f.isLeafNode() {
		return
	}
	var fix *node
	color := f.color()
	if f.left().isLeafNode() {
		fix = f.right()
		t.replace(f, f.right())
	} else if f.right().isLeafNode() {
		fix = f.left()
		t.replace(f, f.left())
	} else {
		successor := t.findSuccessorNode(f.right())
		color = successor.color()
		fix = successor.right()
		if f != successor.parent() {
			t.replace(successor, successor.right())
			successor.setRight(f.right())
		}
		t.replace(f, successor)
		successor.setColor(f.color())
		successor.setLeft(f.left())
	}
	if color == BLACK {
		t.fixRemove(fix)
	}
}

func (t *rbTree) replace(u, v *node) {
	if isRoot(u) {
		t.setRoot(v)
	} else if u.isLeft() {
		u.parent().setLeft(v)
	} else {
		u.parent().setRight(v)
	}
}

func (t *rbTree) fixRemove(n *node) {
START:
	if isRoot(n) || n.isRed() {
		n.turnBlack()
		return
	}
	if n.isLeft() {
		b := n.parent().right()
		if b.isRed() {
			b.turnBlack()
			n.parent().turnRed()
			t.leftRotate(n.parent())
			b = n.parent().right()
			if b.isLeafNode() {
				goto START
			}
		}
		if b.left().isBlack() && b.right().isBlack() {
			b.turnRed()
			n = n.parent()
			goto START
		} else if b.left().isRed() {
			b.left().turnBlack()
			b.turnRed()
			t.rightRotate(b)
			b = n.parent().right()
		}
		if b.right().isRed() {
			b.setColor(b.parent().color())
			b.right().turnBlack()
			b.parent().turnBlack()
			t.leftRotate(b.parent())
			n = t.root()
		}
		goto START
	}
	b := n.parent().left()
	if b.isRed() {
		b.turnBlack()
		n.parent().turnRed()
		t.rightRotate(n.parent())
		b = n.parent().left()
		if b.isLeafNode() {
			goto START
		}
	}
	if b.right().isBlack() && b.left().isBlack() {
		b.turnRed()
		n = n.parent()
		goto START
	} else if b.right().isRed() {
		b.right().turnBlack()
		b.turnRed()
		t.leftRotate(b)
		b = n.parent().left()
	}
	if b.left().isRed() {
		b.setColor(n.parent().color())
		n.parent().turnBlack()
		b.left().turnBlack()
		t.rightRotate(n.parent())
		n = t.root()
	}
}

func (t *rbTree) findSuccessorNode(son *node) *node {
	for son.left().isNonLeafNode() {
		son = son.left()
	}
	return son
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
	if x == nil || x.isLeafNode() || x.right().isLeafNode() {
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
	if x == nil || x.isLeafNode() || x.left().isLeafNode() {
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

func (t *rbTree) setRoot(n *node) bool {
	if n == nil {
		t.r = nil
		return true
	}
	n.p = nil
	t.r = n
	return true
}
