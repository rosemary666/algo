package stack

import (
	"fmt"
	"sync"
)

/*
 * 基于数组实现的栈
 */
type StackOnArray struct {
	data []interface{} //栈的元素列表
	size int           // 栈的元素数目
	mu   sync.Mutex    // 为了并发安全使用的锁
}

func NewStackOnArray() *StackOnArray {
	return &StackOnArray{}
}

func (this *StackOnArray) Push(v interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.data = append(this.data, v)
	this.size++
	return
}

func (this *StackOnArray) Pop() (v interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.IsEmpty() {
		return nil
	}
	v = this.data[this.size-1]
	this.data = this.data[0 : this.size-1]
	this.size--
	return v
}

func (this *StackOnArray) IsEmpty() (ok bool) {
	return this.size == 0
}

func (this *StackOnArray) Top() (v interface{}) {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.Size()-1]
}

func (this *StackOnArray) Reset() {
	this.data = make([]interface{}, 0)
	this.size = 0
	return
}

func (this *StackOnArray) Size() int {
	return this.size
}

func (this *StackOnArray) Print() {
	if this.IsEmpty() {
		fmt.Printf("stack is empty!\n")
		return
	}
	fmt.Printf("stack size is:[%v]\n", this.size)
	fmt.Printf("print stack elem......\n")
	for i := this.size - 1; i >= 0; i-- {
		fmt.Printf("index:%v, elem:[%v]\n", i, this.data[i])
	}
	return
}
