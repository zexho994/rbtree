package red_black_trees

const BLACK = false
const RED = true

type V int

type node struct {
	val   V
	left  *node
	right *node
	color bool
}

func NewNode(v V) *node {
	return &node{
		val:   v,
		left:  nil,
		right: nil,
		color: RED,
	}
}

func (n *node) Val() V {
	return n.val
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) SetLeft(l *node) {
	n.left = l
}

func (n *node) SetRight(r *node) {
	n.right = r
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) IsRed() bool {
	return n.color
}
