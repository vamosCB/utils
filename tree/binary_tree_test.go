package tree

import "testing"

func newTree() *Node {
	root := InitNode(1)
	node2 := InitNode(2)
	node3 := InitNode(3)
	node4 := InitNode(4)
	root.SetLeft(node2)
	root.SetRight(node3)
	node3.SetLeft(node4)
	return root
}

func Test_PreOrder(t *testing.T) {
	testTree := newTree()
	testTree.PostOrder()
}
