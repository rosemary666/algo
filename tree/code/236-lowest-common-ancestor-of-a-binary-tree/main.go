package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 遍历为空或者找到目标点
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	rigth := lowestCommonAncestor(root.Right, p, q)
	// 如果left和rigth都不为空，代表找到祖先节点
	if left != nil && rigth != nil {
		return root
	}
	if left != nil {
		return left
	}
	if rigth != nil {
		return rigth
	}
	return nil
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
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 6}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 0}
	node6 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 4}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node4.Left = node7
	node4.Right = node8
	fmt.Printf("%+v\n", preorderTraversal(lowestCommonAncestor(root, root.Left, root.Right)))

	println("UseCase 2......")
	root = &TreeNode{Val: 3}
	node1 = &TreeNode{Val: 5}
	node2 = &TreeNode{Val: 1}
	node3 = &TreeNode{Val: 6}
	node4 = &TreeNode{Val: 2}
	node5 = &TreeNode{Val: 0}
	node6 = &TreeNode{Val: 8}
	node7 = &TreeNode{Val: 7}
	node8 = &TreeNode{Val: 4}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node4.Left = node7
	node4.Right = node8
	fmt.Printf("%+v\n", preorderTraversal(lowestCommonAncestor(root, root.Left, root.Left.Right.Right)))

	println("UseCase 3......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 2}
	root.Left = node1
	fmt.Printf("%+v\n", preorderTraversal(lowestCommonAncestor(root, root, root.Left)))
}
