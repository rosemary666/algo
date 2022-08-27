package circlelist

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{Data: v, Next: nil}
}

type LinkList struct {
	Head   *LinkNode
	Length uint
}

func NewLinkList() *LinkList {
	return &LinkList{Length: 0}
}

func (this *LinkList) GetHead() *LinkNode {
	return this.Head
}

func (this *LinkList) GetLen() uint {
	return this.Length
}

// 向链表的右边插入
func (this *LinkList) RPush(v interface{}) bool {
	newNode := NewLinkNode(v)
	if this.GetLen() == 0 {
		this.Head = newNode
		newNode.Next = this.Head
		this.Length++
		return true
	}

	curNode := this.Head
	for curNode.Next != this.Head {
		curNode = curNode.Next
	}

	curNode.Next = newNode
	newNode.Next = this.Head
	this.Length++
	return true
}

// 从链表右边开始删除节点
func (this *LinkList) RPop() *LinkNode {
	if this.GetLen() == 0 {
		return nil
	}

	if this.GetHead().Next == this.GetHead() {
		n1 := (*this.GetHead())
		n := &n1
		this.Head = nil
		this.Length--
		return n
	}

	curNode := this.GetHead()
	preNode := this.GetHead()
	for curNode.Next != this.GetHead() {
		preNode = curNode
		curNode = curNode.Next
	}
	preNode.Next = this.Head
	this.Length--
	return curNode
}

func (this *LinkList) Print() {
	if this.GetLen() == 0 {
		return
	}

	curNode := this.GetHead()
	index := 0
	for curNode.Next != this.GetHead() {
		fmt.Printf("[index/curNodeValue] | [%v/%v]\n", index, curNode.Data)
		index++
		curNode = curNode.Next
	}
	fmt.Printf("[index/curNodeValue] | [%v/%v]\n", index, curNode.Data)
}

func (this *LinkList) GetByIndex(index uint) *LinkNode {
	if index >= this.GetLen() {
		return nil
	}
	curNode := this.GetHead()
	for i := 0; i < int(index); i++ {
		curNode = curNode.Next
	}
	return curNode
}
