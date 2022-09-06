package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归遍历
func postorderTraversal(root *TreeNode) []int {
	nodeArr := make([]int, 0)
	traversal(root, &nodeArr)
	return nodeArr
}

func traversal(cur_node *TreeNode, arr *[]int) {
	if cur_node == nil {
		return
	}
	traversal(cur_node.Left, arr)
	traversal(cur_node.Right, arr)
	*arr = append(*arr, cur_node.Val)
}

// 构建二叉树
func buildTree(root *TreeNode, arr []int, i int) {
	if i > len(arr)-1 || len(arr) == 0 {
		return
	}
	root.Val = arr[i]
	i++
	if i <= len(arr)-1 && arr[i] != -1000 {
		root.Left = &TreeNode{}
		buildTree(root.Left, arr, i)
	}
	i++
	if i <= len(arr)-1 && arr[i] != -1000 {
		root.Right = &TreeNode{}
		buildTree(root.Right, arr, i)
	}
}

// 用数据实现一个基本功能的栈
type Stack struct {
	elements []*TreeNode
}

func NewStack() *Stack {
	return &Stack{elements: make([]*TreeNode, 0)}
}

func (s *Stack) IsEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(element *TreeNode) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() *TreeNode {
	if !s.IsEmpty() {
		res := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return res
	}
	return nil
}

// 迭代遍历
func postorderTraversalV1(root *TreeNode) []int {
	nodeArr := make([]int, 0)
	if root == nil {
		return nodeArr
	}
	// 普通栈
	s := NewStack()
	// 输出栈
	out := NewStack()
	// 步骤1: 直到当前结点为空 & 栈空时，循环结束
	for root != nil || !s.IsEmpty() {
		// 步骤2: 判断当前结点是否为空
		if root != nil {
			// 步骤3: 将当前节点输入到普通栈和输出栈
			s.Push(root)
			out.Push(root)
			// 步骤4: 置当前结点的右孩子为当前节点
			root = root.Right
			// 返回步骤1
		} else {
			// 步骤5：出栈栈顶结点
			root = s.Pop()
			// 步骤6：置当前结点的左孩子为当前节点
			root = root.Left
			// 返回步骤1
		}
	}
	for !out.IsEmpty() {
		nodeArr = append(nodeArr, out.Pop().Val)
	}
	return nodeArr
}

func main() {
	// -1000代表空值
	println("UseCase 1......")
	root := &TreeNode{}
	buildTree(root, []int{1, -1000, 2, 3}, 0)

	fmt.Printf("%v\n", postorderTraversal(root))
	fmt.Printf("%v\n", postorderTraversalV1(root))

	println("UseCase 2......")
	root = &TreeNode{}
	buildTree(nil, []int{}, 0)
	fmt.Printf("%v\n", postorderTraversal(nil))
	fmt.Printf("%v\n", postorderTraversalV1(nil))

	println("UseCase 3......")
	root = &TreeNode{}
	buildTree(root, []int{1}, 0)
	fmt.Printf("%v\n", postorderTraversal(root))
	fmt.Printf("%v\n", postorderTraversalV1(root))
}
