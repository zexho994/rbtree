package red_black_trees

const BLACK = false
const RED = true

type V int

type node struct {
	v   V     // val
	l   *node // left node
	r   *node // right node
	p   *node // parent
	c   bool  // color
	nil bool  // leaf-node is nil
}

func newRedNode(v V) *node {
	n := &node{
		v: v,
		c: RED,
	}
	n.l = &node{c: BLACK, nil: true, p: n}
	n.r = &node{c: BLACK, nil: true, p: n}
	return n
}

func newBlackNode(v V) *node {
	n := &node{
		v: v,
		c: BLACK,
	}
	n.l = &node{c: BLACK, nil: true, p: n}
	n.r = &node{c: BLACK, nil: true, p: n}
	return n
}

func newNilNode() *node {
	return &node{
		v:   0,
		c:   BLACK,
		nil: true,
	}
}

func (n *node) val() V {
	return n.v
}

func (n *node) left() *node {
	return n.l
}

func (n *node) right() *node {
	return n.r
}

func (n *node) parent() *node {
	return n.p
}

func (n *node) setParent(p *node) {
	n.p = p
}

func (n *node) brother() *node {
	if n == nil || n.parent() == nil {
		return nil
	}
	if n.isLeft() {
		return n.parent().right()
	} else {
		return n.parent().left()
	}
}

func (n *node) grandfather() *node {
	if n.parent() == nil {
		return nil
	}
	return n.parent().parent()
}

func (n *node) uncle() *node {
	if n.grandfather() == nil {
		return nil
	}
	if n.parent().isLeft() {
		return n.grandfather().right()
	}
	return n.grandfather().left()
}

func (n *node) sonCount() (s int) {
	s = 2
	if n.left().isLeaf() {
		s--
	}
	if n.r.isLeaf() {
		s--
	}
	return
}

func (n *node) isLeafNode() bool {
	return n.nil
}

func (n *node) isNonLeafNode() bool {
	return !n.nil
}

func (n *node) setColor(c bool) {
	n.c = c
}

func (n *node) setLeft(l *node) {
	n.l = l
	l.p = n
}

func (n *node) delLeft() {
	n.setLeft(newNilNode())
}

func (n *node) delRight() {
	n.setRight(newNilNode())
}

func (n *node) setRight(r *node) {
	n.r = r
	r.p = n
}

func (n *node) turnRed() {
	n.c = RED
}

func (n *node) turnBlack() {
	n.c = BLACK
}

func (n *node) findSuccessorNode() (s *node) {
	if n.right().isLeafNode() {
		return nil
	}
	s = n.right()
	for s.left().isLeafNode() {
		s = s.left()
	}
	return
}

func (n *node) isBlack() bool {
	return n.c == BLACK
}

func (n *node) isRed() bool {
	return n.c == RED
}

func (n *node) color() bool {
	return n.isRed()
}

func (n *node) isLeft() bool {
	if n.parent() == nil {
		panic("node p is nil")
	}
	return n.val() < n.parent().val()
}

func (n *node) isRight() bool {
	return !n.isLeft()
}
