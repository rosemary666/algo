package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTravel(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	// 遍历左子树
	inorderTravel(root.Left, arr)
	// 访问根节点
	*arr = append(*arr, root.Val)
	// 遍历右子树
	inorderTravel(root.Right, arr)
}

func getMinimumDifference(root *TreeNode) int {
	// 利用二叉搜索树和中序遍历的特性，可以采用中序遍历，得到一个有序的数组
	arr := make([]int, 0)
	// 中序遍历后得到一个有序的数组
	inorderTravel(root, &arr)
	minValue := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		sub := arr[i] - arr[i-1]
		if sub < 0 {
			sub = -sub
		}
		if sub < minValue {
			minValue = sub
		}
	}
	return minValue
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	fmt.Printf("%+v\n", getMinimumDifference(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 0}
	node2 = &TreeNode{Val: 48}
	node3 = &TreeNode{Val: 12}
	node4 = &TreeNode{Val: 49}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Printf("%+v\n", getMinimumDifference(root))
}
