package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	VirtualHead *Node // 虚拟头节点
}

func Constructor() MyLinkedList {
	node := &Node{}
	node.Next = nil
	return MyLinkedList{VirtualHead: node}
}

// 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	count := 0
	ret := -1
	cur_node := this.VirtualHead.Next
	for index >= 0 && cur_node != nil {
		if index == count {
			ret = cur_node.Val
			break
		}
		count++
		cur_node = cur_node.Next
	}
	return ret
}

// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	cur_head := this.VirtualHead.Next
	// 真实头结点为空
	if cur_head == nil {
		this.VirtualHead.Next = &Node{Val: val}
	} else {
		// 真实头结点非空
		insertNode := &Node{Val: val}
		insertNode.Next = cur_head
		this.VirtualHead.Next = insertNode
	}
	return
}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	cur_node := this.VirtualHead.Next
	pre_node := this.VirtualHead
	for cur_node != nil {
		pre_node = cur_node
		cur_node = cur_node.Next
	}
	pre_node.Next = &Node{Val: val}
}

// 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	count := 0
	cur_node := this.VirtualHead.Next
	pre_node := this.VirtualHead
	for index >= 0 && cur_node != nil {
		if index == count {
			insertNode := &Node{Val: val}
			insertNode.Next = cur_node
			pre_node.Next = insertNode
		}
		count++
		pre_node = cur_node
		cur_node = cur_node.Next
	}
	if index == count {
		this.AddAtTail(val)
	}
}

// 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	count := 0
	cur_node := this.VirtualHead.Next
	pre_node := this.VirtualHead
	for index >= 0 && cur_node != nil {
		if index == count {
			pre_node.Next = cur_node.Next
			break
		}
		count++
		pre_node = cur_node
		cur_node = cur_node.Next
	}
	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func print_linked_list(m MyLinkedList) {
	elements := []int{}
	cur_node := m.VirtualHead.Next
	for cur_node != nil {
		elements = append(elements, cur_node.Val)
		cur_node = cur_node.Next
	}
	fmt.Printf("%v\n", elements)
}

func main() {
	m := Constructor()
	m.AddAtHead(1)
	print_linked_list(m)
	m.AddAtTail(3)
	print_linked_list(m)
	m.AddAtIndex(1, 2)
	print_linked_list(m)
	println(m.Get(1))
	m.DeleteAtIndex(1)
	print_linked_list(m)
	println(m.Get(1))
}
