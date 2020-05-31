package tree

import "fmt"

//Node 结点
type Node struct {
	value interface{}
	left  *Node
	right *Node
}

//InitNode 初始化树根节点
func InitNode(value interface{}) *Node {
	return &Node{value: value}
}

//SetValue ...
func (node *Node) SetValue(value interface{}) {
	if node == nil {
		fmt.Println("this tree is empty!")
		return
	}
	node.value = value
}

//SetRight ...
func (node *Node) SetRight(right *Node) {
	if node == nil {
		fmt.Println("this tree is empty!")
		return
	}
	node.right = right
}

//SetLeft ...
func (node *Node) SetLeft(left *Node) {
	if node == nil {
		fmt.Println("this tree is empty!")
		return
	}
	node.left = left
}

//Print  output node value
func (node *Node) Print() {
	fmt.Print(node.value)
}

//PreOrder 先序遍历
func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.left.PreOrder()
	node.right.PreOrder()
}

//InOrder 中序遍历
func (node *Node) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	node.Print()
	node.right.InOrder()
}

//PostOrder 后序遍历
func (node *Node) PostOrder() {
	if node == nil {
		return
	}
	node.left.PostOrder()
	node.right.PostOrder()
	node.Print()
}

//Depth 树的深度 对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
func (node *Node) Depth() int {
	if node == nil {
		return 0
	}
	leftDepth := node.left.Depth()
	rightDepth := node.right.Depth()
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

//DepthByQueue 深度(非递归)
func (node *Node) DepthByQueue() int {
	if node == nil {
		return 0
	}
	layers := 0
	nodes := []*Node{node}
	for len(nodes) > 0 {
		layers++
		size := len(nodes) //每层的节点数
		count := 0
		for count < size {
			count++
			curNode := nodes[0]
			nodes = nodes[1:]
			if curNode.left != nil {
				nodes = append(nodes, curNode.left)
			}
			if curNode.right != nil {
				nodes = append(nodes, curNode.right)
			}
		}
	}
	return layers
}

//GetLeafNode 获取所有子节点
func (node *Node) GetLeafNode() {
	if node != nil {
		if node.left == nil && node.right == nil {
			fmt.Print(node.value)
		}
		node.left.GetLeafNode()
		node.right.GetLeafNode()
	}
}
