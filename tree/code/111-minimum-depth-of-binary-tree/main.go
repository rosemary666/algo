package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用数组实现个简单的队列功能
type Queue struct {
	elements []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{elements: make([]*TreeNode, 0)}
}

func (q *Queue) Push(node *TreeNode) {
	q.elements = append(q.elements, node)
}

func (q *Queue) Pop() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	node := q.elements[0]
	q.elements = q.elements[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	if q.Size() == 0 {
		return true
	}
	return false
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := NewQueue()
	q.Push(root)
	minDepth := 0
	for !q.IsEmpty() {
		qSize := q.Size() // 代表当前这一层的元素， 一定要先取出size, q.Size()会一直变化
		for i := 0; i < qSize; i++ {
			topNode := q.Pop()
			// 左右子树全部为空，即为叶子节点，找寻终止
			if topNode.Left == nil && topNode.Right == nil {
				minDepth += 1
				return minDepth
			}
			if topNode.Left != nil {
				q.Push(topNode.Left)
			}
			if topNode.Right != nil {
				q.Push(topNode.Right)
			}
		}
		minDepth += 1
	}
	return minDepth
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Printf("%v\n", minDepth(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 2}
	node1 = &TreeNode{Val: 3}
	node2 = &TreeNode{Val: 4}
	node3 = &TreeNode{Val: 5}
	node4 = &TreeNode{Val: 6}
	root.Right = node1
	node1.Right = node2
	node2.Right = node3
	node3.Right = node4
	fmt.Printf("%v\n", minDepth(root))
}
