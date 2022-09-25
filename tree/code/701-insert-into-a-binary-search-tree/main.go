package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
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
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	fmt.Printf("%+v\n", preorderTraversal(insertIntoBST(root, 5)))

	println("UseCase 2......")
	root = &TreeNode{Val: 40}
	node1 = &TreeNode{Val: 20}
	node2 = &TreeNode{Val: 60}
	node3 = &TreeNode{Val: 10}
	node4 = &TreeNode{Val: 30}
	node5 := &TreeNode{Val: 50}
	node6 := &TreeNode{Val: 70}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	fmt.Printf("%+v\n", preorderTraversal(insertIntoBST(root, 25)))
}
