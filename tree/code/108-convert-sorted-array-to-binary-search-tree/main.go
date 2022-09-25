package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 每次二分, 中间节点
	mid := left + (right-left)/2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = travel(nums, left, mid-1)
	root.Right = travel(nums, mid+1, right)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := travel(nums, 0, len(nums)-1)
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
	fmt.Printf("%+v\n", preorderTraversal(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))

	println("UseCase 2......")
	fmt.Printf("%+v\n", preorderTraversal(sortedArrayToBST([]int{1, 3})))
}
