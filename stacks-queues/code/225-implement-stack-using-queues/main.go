package main

import "fmt"

type MyStack struct {
	queue []int // 用数组实现一个队列的功能
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

// 将元素 `x` 压入栈顶。
func (this *MyStack) Push(x int) {
	// 直接把元素加入队列即可
	this.queue = append(this.queue, x)
}

// 移除并返回栈顶元素
func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return 0
	}
	// 弹出除队尾元素的所有元素，并重新追加到队列中
	n := len(this.queue) - 1
	for n != 0 {
		// 模拟队列头弹出元素，并且把队列头元素重新加入队列尾部
		x := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, x)
		n--
	}
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

// 返回栈顶元素
func (this *MyStack) Top() int {
	x := this.Pop()
	if x == 0 {
		return x
	}
	this.Push(x)
	return x
}

// 如果栈是空的，返回 `true` ；否则，返回 `false`
func (this *MyStack) Empty() bool {
	if len(this.queue) != 0 {
		return false
	}
	return true
}

func main() {
	println("UseCase 1......")
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Printf("Peek: %v \n", stack.Top())
	fmt.Printf("Pop: %v \n", stack.Pop())
	fmt.Printf("Empty: %v \n", stack.Empty())

}
