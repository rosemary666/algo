package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, min, max int) bool {
	// 为空，也是有效的
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	// 这里很重要, 所有左子树的值必须小于当前根节点, 右子树的值必须大于当前根节点
	return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	fmt.Printf("%+v\n", isValidBST(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 5}
	node1 = &TreeNode{Val: 1}
	node2 = &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 6}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Printf("%+v\n", isValidBST(root))
}
