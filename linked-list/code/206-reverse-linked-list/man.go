package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法
func reverseList(head *ListNode) *ListNode {
	// 重点是: cur_node初始化是从头结点, 然后pre_node是空指针
	cur_node := head
	var pre_node *ListNode
	for cur_node != nil {
		next_node := cur_node.Next
		cur_node.Next = pre_node
		pre_node = cur_node
		cur_node = next_node
	}
	return pre_node
}

// 递归法
func reverseList1(head *ListNode) *ListNode {
	return reverse_dg(nil, head)
}

func reverse_dg(pre *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverse_dg(cur, next)
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

func print_linked_list(head *ListNode) {
	elements := []int{}
	cur_node := head
	for cur_node != nil {
		elements = append(elements, cur_node.Val)
		cur_node = cur_node.Next
	}
	fmt.Println(elements)
}

func main() {
	println("UseCase 1......")
	h1 := generate_linked_list([]int{1, 2, 3, 4, 5})
	print_linked_list(h1)
	print_linked_list(reverseList(h1))

	println("UseCase 2......")
	h2 := generate_linked_list([]int{1, 2})
	print_linked_list(h2)
	print_linked_list(reverseList(h2))

	println("UseCase 3......")
	h3 := generate_linked_list([]int{})
	print_linked_list(h3)
	print_linked_list(reverseList(h3))
}
