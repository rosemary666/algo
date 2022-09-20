package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode, res *int) {
	// 判断是左叶子节点
	// 这里条件很重要!!!!!, 主要是判断左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res = *res + root.Left.Val
	}
	if root.Left != nil {
		travel(root.Left, res)
	}
	if root.Right != nil {
		travel(root.Right, res)
	}
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	travel(root, &res)
	return res
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
	fmt.Printf("%+v\n", sumOfLeftLeaves(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	fmt.Printf("%+v\n", sumOfLeftLeaves(root))
}
