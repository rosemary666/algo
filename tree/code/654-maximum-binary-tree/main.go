package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 数组为空，终止
	if len(nums) == 0 {
		return nil
	}
	maxIndex := findMaxValIndex(nums)
	// 将数组一分为二, 依次递归
	root := &TreeNode{
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
		Val:   nums[maxIndex],
	}
	return root
}

// 找寻最大值的下标
func findMaxValIndex(nums []int) (index int) {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
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
	nums := []int{3, 2, 1, 6, 0, 5}
	fmt.Printf("%+v\n", preorderTraversal(constructMaximumBinaryTree(nums)))

	println("UseCase 2......")
	nums = []int{3, 2, 1}
	fmt.Printf("%+v\n", preorderTraversal(constructMaximumBinaryTree(nums)))
}
