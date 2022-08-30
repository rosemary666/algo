package singlelist

import "fmt"

type LinkNode struct {
	Data     interface{}
	NextNode *LinkNode
}

func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{Data: v, NextNode: nil}
}

func (this *LinkNode) Next() *LinkNode {
	return this.NextNode
}

func (this *LinkNode) Value() interface{} {
	return this.Data
}

type LinkList struct {
	head   *LinkNode
	length uint
}

func NewLinkList() *LinkList {
	return &LinkList{head: NewLinkNode(nil), length: 0}
}

// 在某个节点后面插入值
func (this *LinkList) InsertAfter(n *LinkNode, v interface{}) (ok bool) {
	if n == nil {
		return false
	}
	newNode := NewLinkNode(v)
	oldNextNode := n.NextNode
	n.NextNode = newNode
	newNode.NextNode = oldNextNode
	this.length++
	return true
}

// 插入到尾部
func (this *LinkList) InsertToTail(v interface{}) (ok bool) {
	node := this.head
	for node.Next() != nil {
		node = node.Next()
	}
	return this.InsertAfter(node, v)
}

// 在某个节点前面插入
func (this *LinkList) InsertBefore(n *LinkNode, v interface{}) (ok bool) {
	if n == this.head || n == nil {
		return false
	}
	curNode := this.head.Next()
	preNode := this.head
	for curNode != nil {
		if curNode == n {
			break
		}
		preNode = curNode
		curNode = curNode.Next()
	}
	if curNode == nil {
		return false
	}
	newNode := NewLinkNode(v)
	preNode.NextNode = newNode
	newNode.NextNode = curNode
	this.length++
	return true
}

// 插入到头部
func (this *LinkList) InsertToHead(v interface{}) (ok bool) {
	return this.InsertAfter(this.head, v)
}

// 返回链表的第一个元素
func (this *LinkList) Front() (n *LinkNode) {
	return this.head.Next()
}

// 返回链表的最后一个元素
func (this *LinkList) Back() (n *LinkNode) {
	curNode := this.head
	for curNode != nil {
		curNode = curNode.Next()
		if curNode.Next() == nil {
			break
		}
	}
	return curNode
}

// 链表长度
func (this *LinkList) Len() (n uint) {
	return this.length
}

// 删除某个节点
func (this *LinkList) Remove(n *LinkNode) (ok bool) {
	if n == nil {
		return false
	}
	preNode := this.head
	curNode := this.head.Next()
	for curNode != nil {
		if n == curNode {
			break
		}
		preNode = curNode
		curNode = curNode.Next()
	}
	if curNode == nil {
		return false
	}
	preNode.NextNode = n.Next()
	n = nil
	this.length--
	return true
}

// 按索引查找某个节点
func (this *LinkList) FindByIndex(index uint) (n *LinkNode) {
	if index >= this.length {
		return nil
	}
	var i = uint(0)
	curNode := this.head.Next()
	for i = 0; i < index; i++ {
		curNode = curNode.Next()
	}
	return curNode
}

func (this *LinkList) Print() {
	curNode := this.head.Next()
	index := 0
	for curNode != nil {
		fmt.Printf("[index/value] is:[%v/%v]\n", index, curNode.Value())
		index++
		curNode = curNode.Next()
	}
	return
}
