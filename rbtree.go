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

}
