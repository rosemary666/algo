package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compareTree(left *TreeNode, right *TreeNode) bool {
	// 判断节点为空的情况
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	// 节点非空，比较值是否相等
	if left.Val != right.Val {
		return false
	}
	return compareTree(left.Left, right.Right) && compareTree(left.Right, right.Left)
}

// 递归法实现
func isSymmetric(root *TreeNode) bool {
	return compareTree(root.Left, root.Right)
}

// 迭代法实现
func isSymmetric1(root *TreeNode) bool {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) != 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		// 节点为空的情况下
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		// 节点非空的情况下
		if left.Val != right.Val {
			return false
		}
		// 重点步骤!!!思路跟递归其实是类似的
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	fmt.Printf("%v\n", isSymmetric1(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 2}
	node3 = &TreeNode{Val: 3}
	node4 = &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Right = node3
	node2.Right = node4
	fmt.Printf("%v\n", isSymmetric1(root))
}
