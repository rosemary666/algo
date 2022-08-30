package stack

import (
	"fmt"
	"sync"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type StackOnLinklist struct {
	top  *LinkNode //栈顶元素
	size int       //栈的元素数目
	mu   sync.Mutex
}

func NewStackOnLinklist() *StackOnLinklist {
	return &StackOnLinklist{
		top:  nil,
		size: 0,
	}
}

func (this *StackOnLinklist) Push(v interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.top == nil {
		node := &LinkNode{data: v}
		this.top = node
	} else {
		preNode := this.top

		newNode := &LinkNode{data: v}
		newNode.next = preNode

		this.top = newNode
	}
	this.size++
	return
}

func (this *StackOnLinklist) Pop() (v interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.IsEmpty() {
		return nil
	}

	v = this.top.data

	this.top = this.top.next
	this.size--
	return v
}

func (this *StackOnLinklist) IsEmpty() (ok bool) {
	return this.size == 0
}

func (this *StackOnLinklist) Top() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	return this.top.data
}

func (this *StackOnLinklist) Reset() {
	this.size = 0
	this.top = nil
	return
}

func (this *StackOnLinklist) Size() int {
	return this.size
}

func (this *StackOnLinklist) Print() {
	if this.IsEmpty() {
		fmt.Printf("stack is empty!\n")
		return
	}
	fmt.Printf("stack size is:[%v]\n", this.size)
	fmt.Printf("print stack elem......\n")

	curNode := this.top
	for curNode != nil {
		fmt.Printf("elem:[%v]\n", curNode.data)
		curNode = curNode.next
	}
	return
}
