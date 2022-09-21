package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 节点为空
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 找到后序数组最后一个元素(就是子树的根节点)在中序数组中的下标, 然后进行二分
	pos := findInorderPos(inorder, postorder[len(postorder)-1])
	// 这里一定要注意边界!!!
	root := &TreeNode{
		Left:  buildTree(inorder[:pos], postorder[:pos]),
		Right: buildTree(inorder[pos+1:], postorder[pos:len(postorder)-1]),
		Val:   inorder[pos],
	}
	return root
}

// 找中序目标值的索引位置
func findInorderPos(inorder []int, target int) int {
	for index, v := range inorder {
		if v == target {
			return index
		}
	}
	return -1
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
	inOrder := []int{9, 3, 15, 20, 7}
	postOrder := []int{9, 15, 7, 20, 3}
	fmt.Printf("preOrder is:%+v\n", preorderTraversal(buildTree(inOrder, postOrder)))

	println("UseCase 2......")
	inOrder = []int{-1}
	postOrder = []int{-1}
	fmt.Printf("preOrder is:%+v\n", preorderTraversal(buildTree(inOrder, postOrder)))
}
