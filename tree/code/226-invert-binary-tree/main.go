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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		qSize := q.Size() // 代表当前这一层的元素， 一定要先取出size, q.Size()会一直变化
		for i := 0; i < qSize; i++ {
			topNode := q.Pop()
			// 交换左右子树
			topNode.Left, topNode.Right = topNode.Right, topNode.Left
			// 将该节点的左右子树分别塞入队列中
			if topNode.Left != nil {
				q.Push(topNode.Left)
			}
			if topNode.Right != nil {
				q.Push(topNode.Right)
			}
		}
	}
	return root
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodeArr := make([][]int, 0)
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		currentLevelNodeArr := make([]int, 0)
		qSize := q.Size() // 代表当前这一层的元素， 一定要先取出size, q.Size()会一直变化
		for i := 0; i < qSize; i++ {
			topNode := q.Pop()
			currentLevelNodeArr = append(currentLevelNodeArr, topNode.Val)
			// 将该节点的左右子树分别塞入队列中
			if topNode.Left != nil {
				q.Push(topNode.Left)
			}
			if topNode.Right != nil {
				q.Push(topNode.Right)
			}
		}
		if len(currentLevelNodeArr) != 0 {
			nodeArr = append(nodeArr, currentLevelNodeArr)
		}
	}
	return nodeArr
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 6}
	node6 := &TreeNode{Val: 9}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	fmt.Printf("raw: %+v\n", levelOrder(root))
	invertTree(root)
	fmt.Printf("after invert: %+v\n", levelOrder(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 2}
	node1 = &TreeNode{Val: 1}
	node2 = &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	fmt.Printf("raw: %+v\n", levelOrder(root))
	invertTree(root)
	fmt.Printf("after invert: %+v\n", levelOrder(root))

	println("UseCase 3......")
	root = nil
	fmt.Printf("raw: %+v\n", levelOrder(root))
	invertTree(root)
	fmt.Printf("after invert: %+v\n", levelOrder(root))

}
