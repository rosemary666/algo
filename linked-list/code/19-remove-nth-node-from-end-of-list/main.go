package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	vitualHead := &ListNode{}
	vitualHead.Next = head

	fastNode := vitualHead
	slowNode := vitualHead
	// fast指针先向前移动n+1位
	for fastNode != nil && n > 0 {
		fastNode = fastNode.Next
		n -= 1
	}

	// fast和slow指针同时向前移动, 直到fast指针遍历到尾指针
	for fastNode != nil {
		// 到最后一个节点了
		if fastNode.Next == nil {
			slowNode.Next = slowNode.Next.Next
			break
		}
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	return vitualHead.Next
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
	println("UseCase 1......")
	h1 := generate_linked_list([]int{1, 2, 3, 4, 5})
	print_linked_list(h1)
	print_linked_list(removeNthFromEnd(h1, 2))
	println("UseCase 2......")
	h2 := generate_linked_list([]int{1})
	print_linked_list(h2)
	print_linked_list(removeNthFromEnd(h2, 1))
	println("UseCase 3......")
	h3 := generate_linked_list([]int{1, 2})
	print_linked_list(h3)
	print_linked_list(removeNthFromEnd(h3, 1))
}
