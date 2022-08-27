package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 先算出headA, headB的长度
	var (
		headALen = 0
		headBLen = 0
		curHeadA = headA
		curHeadB = headB
	)

	for curHeadA != nil {
		headALen += 1
		curHeadA = curHeadA.Next
	}
	for curHeadB != nil {
		headBLen += 1
		curHeadB = curHeadB.Next
	}
	// 计算headA和headB的长度差
	gap := 0
	var longHead, shortHead *ListNode
	if headALen > headBLen {
		gap = headALen - headBLen
		longHead = headA
		shortHead = headB
	} else {
		gap = headBLen - headALen
		longHead = headB
		shortHead = headA
	}

	// 长的链表向前移动gap位
	for gap > 0 {
		longHead = longHead.Next
		gap--
	}
	// 链表长度相等, 开始比较值
	for longHead != nil {
		if longHead == shortHead {
			return longHead
		}
		longHead = longHead.Next
		shortHead = shortHead.Next
	}
	return nil
}

func print_linked_list(head *ListNode) {
	cur_node := head
	elements := []int{}
	for cur_node != nil {
		elements = append(elements, cur_node.Val)
		cur_node = cur_node.Next
	}
	fmt.Printf("%v\n", elements)
}

func generate_linked_list(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	var cur_node = head
	for i, elem := range arr {
		if i == 0 {
			cur_node.Val = elem
		} else {
			cur_node.Next = &ListNode{Val: elem}
			cur_node = cur_node.Next
		}
	}
	return head
}

func main() {
	// 注意: 不是看数值是否相等, 而是比较指针是否相等。
	println("UseCase 1......")
	common := generate_linked_list([]int{8, 4, 5})
	headA := generate_linked_list([]int{4, 1})
	headA.Next = common
	headB := generate_linked_list([]int{5, 0, 1})
	headB.Next = common
	print_linked_list(getIntersectionNode(headA, headB))

	println("UseCase 2......")
	common = generate_linked_list([]int{2, 4})
	headA = generate_linked_list([]int{0, 9, 1})
	headA.Next = common
	headB = generate_linked_list([]int{3})
	headB.Next = common
	print_linked_list(getIntersectionNode(headA, headB))

	println("UseCase 3......")
	headA = generate_linked_list([]int{2, 6, 4})
	headB = generate_linked_list([]int{1, 5})
	print_linked_list(getIntersectionNode(headA, headB))
}
