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

func (n *node) Brother() *node {
	if n == nil || n.Parent() == nil {
		return nil
	}
	if n.IsLeft() {
		return n.Parent().Right()
	} else {
		return n.Parent().Left()
	}
}

func (n *node) Grandfather() *node {
	if n.Parent() == nil {
		return nil
	}
	return n.Parent().Parent()
}

func (n *node) Uncle() *node {
	if n.Grandfather() == nil {
		return nil
	}
	if n.Parent().IsLeft() {
		return n.Grandfather().Right()
	}
	return n.Grandfather().Left()
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
