package main

const BLACK = false
const RED = true
type T int

type node struct{
	val T 
	left *node
	right *node
	color bool
}

func (n *node) Val() T{
	return n.val
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node{
	return n.right
}

func (n *node) IsRed() bool {
	return n.color
}