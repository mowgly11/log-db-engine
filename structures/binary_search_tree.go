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

func Search(root *BinaryTreeNode, value int) bool {
	if root == nil {
		return false
	}
	if value == root.Value {
		return true
	}
	if value < root.Value {
		return Search(root.Left, value)
	}
	return Search(root.Right, value)
}