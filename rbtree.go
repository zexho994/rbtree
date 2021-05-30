package red_black_trees

type rbTree struct {
	root *node
}

func NewRbTree(v V) *rbTree {
	return &rbTree{
		root: NewNode(v),
	}
}

func (t *rbTree) Root() *node {
	return t.root
}

func (t *rbTree) Find(v V) *node {
	m := t.root
	for m != nil {
		if v > m.val {
			m = m.right
		} else if v < m.val {
			m = m.left
		} else {
			return m
		}
	}

	return nil
}
