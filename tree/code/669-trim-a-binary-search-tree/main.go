package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	// 寻找符合区间[low, high]的节点
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 寻找符合区间[low, high]的节点
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)   // root->left接入符合条件的左孩子
	root.Right = trimBST(root.Right, low, high) // root->left接入符合条件的右孩子
	return root
}

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
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: 2}
	root.Left = node1
	root.Right = node2
	fmt.Printf("%+v\n", preorderTraversal(trimBST(root, 1, 2)))

	println("UseCase 2......")
	root = &TreeNode{Val: 3}
	node1 = &TreeNode{Val: 0}
	node2 = &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Right = node3
	node3.Left = node4
	fmt.Printf("%+v\n", preorderTraversal(trimBST(root, 1, 3)))
}
