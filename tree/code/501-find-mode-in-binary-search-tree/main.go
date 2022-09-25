package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var (
		res      = make([]int, 0)
		maxCount = 0 // 最大统计数
		count    = 0
		pre      *TreeNode
		travel   func(root *TreeNode)
	)
	travel = func(root *TreeNode) {
		// 二叉搜索树通过中序遍历能得到有序数组
		if root == nil {
			return
		}
		// 遍历左子树
		travel(root.Left)
		// 处理根节点
		if pre == nil { // 第一个节点
			count = 1
		} else if root.Val == pre.Val { // 与前一个节点数值相同
			count += 1
		} else {
			count = 1 // 与前一个节点数值不同
		}
		pre = root             // 更新上一个节点
		if count == maxCount { // 要考虑一下存在相同的情况, 比如有两个值，一个1，一个2， 这个时候应该返回[1, 2]
			res = append(res, root.Val)
		}
		if count > maxCount { //大于记录的最大值, 这个时候res数组需要重置
			maxCount = count
			res = make([]int, 0) //!!!!数组重置
			res = append(res, root.Val)
		}
		// 遍历右子树
		travel(root.Right)
	}
	travel(root)
	return res
}

func main() {
	println("UseCase 1......")
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 2}
	root.Right = node1
	node1.Left = node2
	fmt.Printf("%+v\n", findMode(root))

	println("UseCase 2......")
	root = &TreeNode{Val: 0}
	fmt.Printf("%+v\n", findMode(root))
}
