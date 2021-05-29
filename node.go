package main

const BLACK = false
const RED = true

type node struct{
	left *node
	right *node
	color bool
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