package red_black_trees

type rbTree struct {
	root *node
}

func NewRbTree(v V) *rbTree {
	return &rbTree{
		root: NewNode(v),
	}
}

func (tree *rbTree) Find(v V) *node {
	m := tree.root
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
