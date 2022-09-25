package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到目标树
	if root.Val == val {
		return root
	}

	// 左子树非空，并且左子树根节点大于目标值(左子树值均小于根节点值)
	if root.Left != nil && root.Val > val {
		return searchBST(root.Left, val)
	}
	// 右子树非空，并且右子树根节点小于于目标值(右子树值均大于根节点值)
	if root.Right != nil && root.Val < val {
		return searchBST(root.Right, val)
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
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	fmt.Printf("%+v\n", preorderTraversal(searchBST(root, 2)))

	println("UseCase 2......")
	root = &TreeNode{Val: 4}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 7}
	node3 = &TreeNode{Val: 1}
	node4 = &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	fmt.Printf("%+v\n", preorderTraversal(searchBST(root, 5)))
	return
}
