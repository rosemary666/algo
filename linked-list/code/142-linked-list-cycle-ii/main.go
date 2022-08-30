package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 定义一快一慢两个指针
	var (
		fast = head
		slow = head
	)
	// 找环
	// 快指针每次移动两位, 慢指针每次移动一位
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// fast和slow相遇, 证明有环
		if fast == slow {
			// 寻找环的位置, 头节点和slow节点每次都向前移动一位，直到相遇，即该位置就是环所在的位置
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
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
	head := generate_linked_list([]int{3, 2, 0, -4})
	head.Next.Next.Next = head.Next //构造一个环(-4指向2)
	dest_node := detectCycle(head)
	if dest_node != nil {
		println(detectCycle(head).Val)
	} else {
		println("no cycle...")
	}

	println("UseCase 2......")
	head = generate_linked_list([]int{1, 2})
	head.Next = head //构造一个环(2指向1)
	dest_node = detectCycle(head)
	if dest_node != nil {
		println(detectCycle(head).Val)
	} else {
		println("no cycle...")
	}

	println("UseCase 3......")
	head = generate_linked_list([]int{1})
	dest_node = detectCycle(head)
	if dest_node != nil {
		println(detectCycle(head).Val)
	} else {
		println("no cycle...")
	}

}
