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
	n := &node{
		val:   v,
		color: RED,
		isNil: false,
	}
	n.left = &node{color: BLACK, isNil: true, parent: n}
	n.right = &node{color: BLACK, isNil: true, parent: n}
	return n
}

func NewBlackNode(v V) *node {
	n := &node{
		val:   v,
		color: BLACK,
		isNil: false,
	}
	n.left = &node{color: BLACK, isNil: true, parent: n}
	n.right = &node{color: BLACK, isNil: true, parent: n}
	return n
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

func (n *node) TurnRed() {
	n.color = RED
}

func (n *node) TurnBlack() {
	n.color = BLACK
}

func (n *node) IsBlack() bool {
	return n.color == BLACK
}

func (n *node) IsRed() bool {
	return n.color == RED
}

func (n *node) IsLeft() bool {
	if n.Parent() == nil {
		panic("node parent is nil")
	}
	if n.Val() < n.Parent().Val() {
		return true
	} else {
		return false
	}
}

func (n *node) IsRight() bool {
	return !n.IsLeft()
}
