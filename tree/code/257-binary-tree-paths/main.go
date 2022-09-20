package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode, res *[]string, s string) {
	// 如果是叶子结果, 添加路径到数组中
	if root.Left == nil && root.Right == nil {
		*res = append(*res, s+strconv.Itoa(root.Val))
		return
	}
	// 本质其实就是前序遍历
	s = s + strconv.Itoa(root.Val) + "->"
	if root.Left != nil {
		travel(root.Left, res, s)
	}
	if root.Right != nil {
		travel(root.Right, res, s)
	}
}

// 递归法实现
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	travel(root, &res, "")
	return res
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 5}
	root.Left = node1
	root.Right = node2
	node1.Right = node3
	fmt.Printf("%+v\n", binaryTreePaths(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	fmt.Printf("%+v\n", binaryTreePaths(root))
}
