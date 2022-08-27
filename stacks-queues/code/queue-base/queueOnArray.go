package queue

import "fmt"

/*
 * 基于数组实现队列
 */
type QueueOnArray struct {
	q        []interface{} // 队列
	capacity int           // 容量
	head     int           // 队列头
	tail     int           // 队列尾
}

func NewQueueOnArray(n int) *QueueOnArray {
	if n <= 0 {
		panic("queue capacity <= 0")
	}
	return &QueueOnArray{q: make([]interface{}, n), capacity: n, head: 0, tail: 0}
}

func (this *QueueOnArray) EnQueue(v interface{}) (ok bool) {
	// 队列已满
	if this.IsFull() {
		return false
	}
	/**
	 * 如果tail = capacity，但是head != 0,说明前有数据删除，队列未满，需要数据迁移
	 */
	if this.tail == this.capacity && this.head != 0 {
		for i := this.head; i < this.capacity; i++ {
			this.q[i-this.head] = this.q[i]
		}
		this.tail = this.tail - this.head
		this.head = 0
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *QueueOnArray) DeQueue() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	v = this.q[this.head]
	// 第一个元素后移一次，这样做的好处是在出队时不需要数据迁移
	this.head++
	return
}

func (this *QueueOnArray) IsFull() (ok bool) {
	/*
	 * 判断队列满了的条件，tail = capacity,head = 0,
	 */
	return this.tail == this.capacity && this.head == 0
}

func (this *QueueOnArray) IsEmpty() (ok bool) {
	return this.head == this.tail
}

func (this *QueueOnArray) Peek() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	return this.q[this.head]
}

func (this *QueueOnArray) Reset() {
	this.q = make([]interface{}, 0, this.capacity)
	this.tail = 0
	this.head = 0
	return
}

func (this *QueueOnArray) Size() int {
	return this.tail - this.head
}

func (this *QueueOnArray) Print() {
	if this.IsEmpty() {
		fmt.Printf("empty queue!\n")
		return
	}
	msg := "head"
	for i := this.head; i < this.tail; i++ {
		msg += fmt.Sprintf(">-%+v", this.q[i])
	}
	msg += ">-tail"
	fmt.Println(msg)
	return
}
