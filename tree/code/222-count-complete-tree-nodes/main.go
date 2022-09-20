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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := NewQueue()
	q.Push(root)
	nodeSum := 0
	for !q.IsEmpty() {
		qSize := q.Size() // 代表当前这一层的元素， 一定要先取出size, q.Size()会一直变化
		nodeSum += qSize
		for i := 0; i < qSize; i++ {
			topNode := q.Pop()
			if topNode.Left != nil {
				q.Push(topNode.Left)
			}
			if topNode.Right != nil {
				q.Push(topNode.Right)
			}
		}
	}
	return nodeSum
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 4}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 6}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	fmt.Printf("%v\n", countNodes(root))

	println("UseCase 2......")
	root = nil
	fmt.Printf("%v\n", countNodes(root))

	println("UseCase 3......")
	root = &TreeNode{Val: 1}
	fmt.Printf("%v\n", countNodes(root))
}
