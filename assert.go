package red_black_trees

func NotNull(n *node) {
	if n == nil {
		panic("fail")
	}
}

func IsTrue(c bool) {
	if !c {
		panic("fail")
	}
}

func NotTrue(c bool) {
	if c {
		panic("fail")
	}
}
