package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode, sum *int) {
	// 为空，结束该次递归，开始返回
	if root == nil {
		return
	}
	// 针对二叉搜索树，按照右、中、左的顺序遍历, 即可实现累加
	// 遍历右子树
	travel(root.Right, sum)
	// 处理根节点
	*sum = *sum + root.Val
	root.Val = *sum
	// 遍历左子树
	travel(root.Left, sum)
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int = 0
	travel(root, &sum)
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
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 3}
	node8 := &TreeNode{Val: 8}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node4.Right = node7
	node2.Left = node5
	node2.Right = node6
	node6.Right = node8
	fmt.Printf("%+v\n", preorderTraversal(convertBST(root)))

	println("UseCase 2......")
	root = &TreeNode{Val: 0}
	node1 = &TreeNode{Val: 1}
	root.Right = node1
	fmt.Printf("%+v\n", preorderTraversal(convertBST(root)))

	println("UseCase 3......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 0}
	node2 = &TreeNode{Val: 2}
	root.Left = node1
	root.Right = node2
	fmt.Printf("%+v\n", preorderTraversal(convertBST(root)))

	println("UseCase 4......")
	root = &TreeNode{Val: 3}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 4}
	node3 = &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	fmt.Printf("%+v\n", preorderTraversal(convertBST(root)))
}
