package doublelist

import "fmt"

type LinkNode struct {
	Next *LinkNode
	Pre  *LinkNode
	Data interface{}
}

func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{Data: v, Next: nil, Pre: nil}
}

type LinkList struct {
	Head   *LinkNode
	Tail   *LinkNode
	length uint
}

func NewLinkList() *LinkList {
	return &LinkList{
		Head:   nil,
		Tail:   nil,
		length: 0,
	}
}

func (this *LinkList) Len() (length uint) {
	return this.length
}

func (this *LinkList) GetHead() *LinkNode {
	return this.Head
}

func (this *LinkList) GetTail() *LinkNode {
	return this.Tail
}

// 插入到尾部
func (this *LinkList) RPush(v interface{}) bool {
	newNode := NewLinkNode(v)
	if this.Len() == 0 {
		this.Tail = newNode
		this.Head = newNode
	} else {
		tail := this.Tail
		tail.Next = newNode
		newNode.Pre = tail
		this.Tail = newNode
	}
	this.length++
	return true
}

// 插入到头部
func (this *LinkList) LPush(v interface{}) bool {
	newNode := NewLinkNode(v)
	if this.Len() == 0 {
		this.Tail = newNode
		this.Head = newNode
	} else {
		head := this.Head
		head.Pre = newNode
		newNode.Next = head
		this.Head = newNode
	}
	this.length++
	return true
}

// 从链表左边取出一个元素
func (this *LinkList) LPop() *LinkNode {
	if this.Len() == 0 {
		return nil
	}
	popNode := this.Head

	if popNode.Next == nil {
		//链表为空
		this.Head = nil
		this.Tail = nil
	} else {
		this.Head = popNode.Next
	}
	this.length--
	return popNode
}

// 从链表右边取出一个元素
func (this *LinkList) RPop() *LinkNode {
	if this.Len() == 0 {
		return nil
	}
	popNode := this.Tail

	if popNode.Pre == nil {
		//链表为空
		this.Head = nil
		this.Tail = nil
	} else {
		this.Tail = popNode.Pre
	}
	this.length--
	return popNode
}

// 按索引查找某个节点
func (this *LinkList) FindByIndex(index uint) (n *LinkNode) {
	if index >= this.length {
		return nil
	}
	var i = uint(0)
	curNode := this.Head
	for i = 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode
}

// 链表打印
func (this *LinkList) Print() {
	curNode := this.Head
	index := 0
	if this.Len() == 0 {
		return
	}
	for curNode.Next != nil {
		if curNode.Pre != nil {
			fmt.Printf("[index/curNodeValue|preNodeValue|nextNodeValue] is:[%v/%v|%v|%v]\n", index, curNode.Data,
				curNode.Pre.Data, curNode.Next.Data)
		} else {
			fmt.Printf("[index/curNodeValue|preNodeValue|nextNodeValue] is:[%v/%v|null|%v\n", index, curNode.Data, curNode.Next.Data)
		}
		index++
		curNode = curNode.Next
	}
	if curNode.Pre != nil {
		fmt.Printf("[index/curNodeValue|preNodeValue|nextNodeValue] is:[%v/%v|%v|null\n", index, curNode.Data, curNode.Pre.Data)
	} else {
		fmt.Printf("[index/curNodeValue|preNodeValue|nextNodeValue] is:[%v/%v|null|null\n", index, curNode.Data)
	}
}
