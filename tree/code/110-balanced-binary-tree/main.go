package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(max(maxDepth(root.Left), maxDepth(root.Right)) + 1)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return true
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Printf("%v\n", isBalanced(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 2}
	node3 = &TreeNode{Val: 3}
	node4 = &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 4}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node3.Left = node5
	node3.Right = node6
	fmt.Printf("%v\n", isBalanced(root))

	println("UseCase 3......")
	fmt.Printf("%v\n", isBalanced(nil))
}
