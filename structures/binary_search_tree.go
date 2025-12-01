package structures

type BinaryTreeNode struct {
	Value int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

func Insert(root *BinaryTreeNode, value int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{Value: value}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return root
}