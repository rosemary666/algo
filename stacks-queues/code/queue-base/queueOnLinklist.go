package queue

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type QueueOnLinklist struct {
	capacity int       // 容量
	head     *LinkNode //对头元素
	tail     *LinkNode //队尾元素
	size     int       // 队列大小
}

func NewQueueOnLinklist(n int) *QueueOnLinklist {
	if n <= 0 {
		panic("queue capacity <= 0")
	}
	q := &QueueOnLinklist{
		capacity: n,
		head:     nil,
		tail:     nil,
		size:     0,
	}
	return q
}

func (this *QueueOnLinklist) EnQueue(v interface{}) (ok bool) {
	if this.IsFull() {
		return false
	}
	node := &LinkNode{data: v}
	if this.IsEmpty() {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.size++
	return true
}

func (this *QueueOnLinklist) DeQueue() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	v = this.head.data
	this.head = this.head.next
	this.size--
	return v
}

func (this *QueueOnLinklist) IsFull() (ok bool) {
	return this.size == this.capacity
}

func (this *QueueOnLinklist) IsEmpty() (ok bool) {
	return this.size == 0
}

func (this *QueueOnLinklist) Peek() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	return this.head.data
}

func (this *QueueOnLinklist) Reset() {
	this.head = nil
	this.tail = nil
	this.size = 0
	return
}

func (this *QueueOnLinklist) Size() int {
	return this.size
}

func (this *QueueOnLinklist) Print() {
	if this.IsEmpty() {
		fmt.Println("empty queue!")
		return
	}
	curNode := this.head
	msg := "head"
	for curNode != nil {
		msg += fmt.Sprintf("->%+v", curNode.data)
		curNode = curNode.next
	}
	msg += "->tail"
	fmt.Println(msg)
	return
}
