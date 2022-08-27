package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var vitualHead = &ListNode{} //定义一个虚拟节点
	vitualHead.Next = head
	cur_node := vitualHead
	for cur_node.Next != nil && cur_node.Next.Next != nil { //终止条件很关键，以奇数个和偶数个作为判断条件, 且cur_node.Next需要写在第一位, 防止cur_node.Next.Next可能越界
		temp := cur_node.Next
		temp1 := cur_node.Next.Next.Next
		cur_node.Next = cur_node.Next.Next // 第一步
		cur_node.Next.Next = temp          // 第二步
		temp.Next = temp1                  // 第三步
		cur_node = cur_node.Next.Next      // cur_node向前移动两位
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
	h1 := generate_linked_list([]int{1, 2, 3, 4})
	print_linked_list(h1)
	print_linked_list(swapPairs(h1))

	println("UseCase 2......")
	h2 := generate_linked_list([]int{})
	print_linked_list(h2)
	print_linked_list(swapPairs(h2))

	println("UseCase 3......")
	h3 := generate_linked_list([]int{1})
	print_linked_list(h3)
	print_linked_list(swapPairs(h3))
}
