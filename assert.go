package rbtree

func AssertNotNull(n *node) {
	if n == nil {
		panic("fail")
	}
}

func AssertIsTrue(c bool) {
	if !c {
		panic("fail")
	}
}

func AssertNotTrue(c bool) {
	if c {
		panic("fail")
	}
}
