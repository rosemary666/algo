package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	virtualHead := &ListNode{}
	virtualHead.Next = head
	cur_node := virtualHead
	for cur_node != nil && cur_node.Next != nil {
		if cur_node.Next.Val == val {
			cur_node.Next = cur_node.Next.Next
		} else {
			cur_node = cur_node.Next
		}
	}
	return virtualHead.Next
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
	head1 := generate_linked_list([]int{1, 2, 6, 3, 4, 5, 6})
	print_linked_list(head1)
	print_linked_list(removeElements(head1, 6))

	println("UseCase 2......")
	head2 := generate_linked_list([]int{})
	print_linked_list(head2)
	print_linked_list(removeElements(head2, 1))

	println("UseCase 3......")
	head3 := generate_linked_list([]int{7, 7, 7, 7})
	print_linked_list(head3)
	print_linked_list(removeElements(head3, 7))
}
