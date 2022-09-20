package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}
	// 是叶子节点
	if root.Left == nil && root.Right == nil {
		sum += root.Val
		if sum == targetSum {
			return true
		}
		return false
	}
	sum += root.Val
	return travel(root.Left, targetSum, sum) || travel(root.Right, targetSum, sum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	return travel(root, targetSum, sum)
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 8}
	node3 := &TreeNode{Val: 11}
	node4 := &TreeNode{Val: 13}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Right = node8
	fmt.Printf("%+v\n", hasPathSum(root, 22))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	fmt.Printf("%+v\n", hasPathSum(root, 5))

	println("UseCase 3......")
	root = nil
	fmt.Printf("%+v\n", hasPathSum(root, 0))
}
