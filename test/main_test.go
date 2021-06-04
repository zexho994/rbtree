package test

import (
	"fmt"
	"github.com/zexho994/rbtree"
	"testing"
)

func Test_Using(t *testing.T) {
	// create a rb-tree
	tree := rbtree.NewRbTree(100)

	tree.Insert(10)
	tree.Insert(120)

	n1 := tree.Find(120)
	fmt.Println(n1)

	tree.Remove(10)
}
