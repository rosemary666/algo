package queue

import (
	"fmt"
)

/*
 * 基于数组实现循环队列
 */
type CircularQueue struct {
	q        []interface{} // 队列
	capacity int           // 容量
	head     int           // 队列头
	tail     int           // 队列尾
	size     int           // 队列实际大小
}

func NewCircularQueue(n int) *CircularQueue {
	if n <= 0 {
		panic("queue capacity <= 0")
	}
	return &CircularQueue{q: make([]interface{}, n), capacity: n, head: 0, tail: 0, size: 0}
}

// 将牺牲一个存储空间,否则无法区分队列是满还是空
func (this *CircularQueue) EnQueue(v interface{}) (ok bool) {
	// 队列已满
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	this.size++
	return true
}

func (this *CircularQueue) DeQueue() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	v = this.q[this.head]
	// 第一个元素后移一次，这样做的好处是在出队时不需要数据迁移
	this.head = (this.head + 1) % this.capacity
	this.size--
	return
}

// 队列满的条件: (tail+1) % capacity = head
func (this *CircularQueue) IsFull() (ok bool) {
	return (this.tail+1)%this.capacity == this.head
}

// 队列空条件: head==tail
func (this *CircularQueue) IsEmpty() (ok bool) {
	return this.head == this.tail
}

func (this *CircularQueue) Peek() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	return this.q[this.head]
}

func (this *CircularQueue) Reset() {
	this.q = make([]interface{}, 0, this.capacity)
	this.tail = 0
	this.head = 0
	this.size = 0
	return
}

func (this *CircularQueue) Size() int {
	return this.size
}

func (this *CircularQueue) Print() {
	if this.IsEmpty() {
		fmt.Printf("empty queue!\n")
		return
	}
	msg := "head"
	index := this.head
	for {
		msg += fmt.Sprintf(">-%+v", this.q[index])
		index = (index + 1) % this.capacity
		if index == this.tail {
			break
		}
	}
	msg += ">-tail"
	fmt.Println(msg)
	return
}
