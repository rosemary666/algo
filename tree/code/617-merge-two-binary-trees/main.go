package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 只有一边的情况下, 直接返回另一边
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root := &TreeNode{
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
		Val:   root1.Val + root2.Val, // 节点值相加
	}
	return root
}

// 递归遍历
func preorderTraversal(root *TreeNode) []int {
	nodeArr := make([]int, 0)
	traversal(root, &nodeArr)
	return nodeArr
}

func traversal(cur_node *TreeNode, arr *[]int) {
	if cur_node == nil {
		return
	}
	*arr = append(*arr, cur_node.Val)
	// fmt.Printf("%+v\n", cur_node.Val)
	traversal(cur_node.Left, arr)
	traversal(cur_node.Right, arr)
}

func main() {
	println("UseCase 1......")
	root1 := &TreeNode{Val: 1}
	node11 := &TreeNode{Val: 3}
	node12 := &TreeNode{Val: 2}
	node13 := &TreeNode{Val: 5}
	root1.Left = node11
	root1.Right = node12
	node11.Left = node13

	root2 := &TreeNode{Val: 2}
	node21 := &TreeNode{Val: 1}
	node22 := &TreeNode{Val: 3}
	node23 := &TreeNode{Val: 4}
	node24 := &TreeNode{Val: 7}
	root2.Left = node21
	root2.Right = node22
	node11.Left = node13
	node21.Right = node23
	node22.Right = node24
	fmt.Printf("%+v\n", preorderTraversal(mergeTrees(root1, root2)))

	println("UseCase 2......")
	root1 = &TreeNode{Val: 1}
	root2 = &TreeNode{Val: 1}
	node21 = &TreeNode{Val: 2}
	root2.Left = node21
	fmt.Printf("%+v\n", preorderTraversal(mergeTrees(root1, root2)))
}
