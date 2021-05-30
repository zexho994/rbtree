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

// 如果存在值为v的节点，返回该节点，该节点肯定不是叶子节点
// 如果不存在，返回的节点 IsLeaf() == true
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
	parent := node.Parent()
	if v < parent.Val() {
		parent.SetLeft(NewRedNode(v))
	} else {
		parent.SetRight(NewRedNode(v))
	}

CASE1:
	if IsRoot(node) || node.Parent().IsBlack() {
		return
	}

	if node.Uncle() != nil && node.Uncle().IsRed() {
		node.Parent().TurnBlack()
		node.Uncle().TurnBlack()
		if IsRoot(node.Grandfather()) {
			return
		}
		node.Grandfather().TurnRed()
		node = node.Grandfather()
		goto CASE1
	}

	if !((node.IsLeft() && node.Parent().IsLeft()) || (node.IsRight() && node.Parent().IsRight())) {
		if node.IsLeft() {

		} else {

		}
	}

	node.Parent().TurnRed()
	node.Grandfather().TurnBlack()
	if node.IsLeft() {

	} else {

	}

}

func IsRoot(n *node) bool {
	return n.parent == nil
}

func leftRotate(x *node) bool {
	var c *node
	if c = x.Right(); c.IsLeaf() {
		return false
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
