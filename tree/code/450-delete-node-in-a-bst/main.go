package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 没找到
	if root == nil {
		return root
	}
	// 找到目标值
	if root.Val == key {
		// 节点的左右子树均为空，即该节点为叶子节点，舍弃即可
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 左子树为空，右子树非空, 右子树替换该节点
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		// 左子树非空，右子树为空，左子树替换该节点
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// 左右子树均非空
		if root.Left != nil && root.Right != nil {
			// 找到右子树的最左叶子节点, 将左子树作为最左叶子节点的左子树插入
			tmp := root.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			tmp.Left = root.Left
			// 原节点左子树置空
			root.Left = nil
			// 原节点用右子树替换
			root = root.Right
			return root
		}
	}
	// 寻找左子树
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
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
	root := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Right = node5
	fmt.Printf("%+v\n", preorderTraversal(deleteNode(root, 3)))

	println("UseCase 2......")
	root = &TreeNode{Val: 5}
	node1 = &TreeNode{Val: 3}
	node2 = &TreeNode{Val: 6}
	node3 = &TreeNode{Val: 2}
	node4 = &TreeNode{Val: 4}
	node5 = &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Right = node5
	fmt.Printf("%+v\n", preorderTraversal(deleteNode(root, 0)))

	println("UseCase 3......")
	root = nil
	fmt.Printf("%+v\n", preorderTraversal(deleteNode(root, 0)))
}
