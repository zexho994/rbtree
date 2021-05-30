package red_black_trees

// rb.node is do not allow duplication
type rbTree struct {
	root *node
}

func NewRbTree(v V) *rbTree {
	return &rbTree{
		root: NewBlackNode(v),
	}
}

func (t *rbTree) Root() *node {
	return t.root
}

// Find 值为v的节点
// 如果存在，返回该节点，该节点肯定不是叶子节点
// 否则，返回的节点 IsLeaf() == true
func (t *rbTree) Find(v V) *node {
	m := t.Root()
	for m.IsNonLeaf() {
		if v > m.Val() {
			m = m.Right()
		} else if v < m.Val() {
			m = m.Left()
		} else {
			return m
		}
	}

	return m
}

func (t *rbTree) Insert(v V) {
	node := t.Find(v)
	if node.IsNonLeaf() {
		return
	}
	// 保存新的节点到对应的位置
	if parent := node.Parent(); v < parent.Val() {
		parent.SetLeft(NewRedNode(v))
	} else {
		parent.SetRight(NewRedNode(v))
	}

CASE1:
	if node.Parent().IsBlack() {
		return
	}
	if node.Uncle() == nil || node.Uncle().IsBlack() {
		goto CASE2
	}

	node.Parent().TurnBlack()
	node.Uncle().TurnBlack()
	if IsRoot(node.Grandfather()) {
		return
	}
	node.Grandfather().TurnRed()
	node = node.Grandfather()
	goto CASE1

CASE2: // 当 node父节点的关系 != node父节点与祖父节点的关系 时成立
	if (node.IsLeft() && node.Parent().IsLeft()) || (node.IsRight() && node.Parent().IsRight()) {
		goto CASE3
	}
	if node.IsLeft() {
		node = node.parent
		t.rightRotate(node)
	} else {
		node = node.parent
		t.leftRotate(node)
	}

CASE3: // 当 node父节点的关系 == node父节点与祖父节点的关系 时成立
	if node.IsLeft() {
		t.rightRotate(node.Grandfather())
	} else {
		t.leftRotate(node.Grandfather())
	}
	node.Parent().TurnBlack()
	node.Brother().TurnRed()
}

func IsRoot(n *node) bool {
	if n == nil {
		return false
	}
	return n.parent == nil
}

// 1. x父指针的子指针指向c
// 2. x的右子指针指向c的左子节点
// 3. c的左子指针指向x
func (t *rbTree) leftRotate(x *node) bool {
	if x == nil || x.IsLeaf() || x.Right().IsLeaf() {
		return false
	}
	c := x.Right()
	if t.Root() == x {
		t.setRoot(c)
		x.SetRight(c.Left())
		c.SetLeft(x)
		return true
	}

	if x.IsLeft() {
		x.Parent().SetLeft(c)
	} else {
		x.Parent().SetRight(c)
	}
	x.SetRight(c.Left())
	c.SetLeft(x)
	return true
}

// 1.x父节点的子指针指向x左节点B
// 2.x左节点指向B右节点
// 3.B右节点指向x
func (t *rbTree) rightRotate(x *node) bool {
	if x == nil || x.IsLeaf() || x.Left().IsLeaf() {
		return false
	}
	b := x.Left()
	if t.Root() == x {
		t.setRoot(b)
		x.SetLeft(b.Right())
		b.SetRight(x)
		return true
	}
	if x.IsLeft() {
		x.Parent().SetLeft(b)
	} else {
		x.Parent().SetRight(b)
	}
	x.SetLeft(b.Right())
	b.SetRight(x)
	return true
}

func (t *rbTree) setRoot(n *node) {
	if n == nil {
		return
	}
	n.parent = nil
	t.root = n
}
