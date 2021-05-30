package red_black_trees

const BLACK = false
const RED = true

type V int

type node struct {
	val    V
	left   *node
	right  *node
	parent *node
	color  bool
	isNil  bool // leaf-node is nil
}

func NewRedNode(v V) *node {
	return &node{
		val:   v,
		left:  &node{color: BLACK, isNil: true},
		right: &node{color: BLACK, isNil: true},
		color: RED,
		isNil: false,
	}
}

func NewBlackNode(v V) *node {
	return &node{
		val:   v,
		left:  &node{color: BLACK, isNil: true},
		right: &node{color: BLACK, isNil: true},
		color: BLACK,
		isNil: false,
	}
}

func (n *node) Val() V {
	return n.val
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) Parent() *node {
	return n.parent
}

func (n *node) IsLeaf() bool {
	return n.isNil
}

func (n *node) IsNonLeaf() bool {
	return !n.isNil
}

func (n *node) SetLeft(l *node) {
	n.left = l
	l.parent = n
}

func (n *node) SetRight(r *node) {
	n.right = r
	r.parent = n
}

func (n *node) IsRed() bool {
	return n.color
}
