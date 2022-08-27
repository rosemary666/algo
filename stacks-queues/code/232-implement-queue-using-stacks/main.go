package main

import "fmt"

type MyQueue struct {
	stackIn  []int // 输入栈
	stackOut []int // 输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// 将元素 `x` 推到队列的末尾
func (this *MyQueue) Push(x int) {
	// 直接像输入栈塞入数据即可
	this.stackIn = append(this.stackIn, x)
	return
}

// 从队列的开头移除并返回元素
func (this *MyQueue) Pop() int {
	x := 0
	// 如果输出栈为空, 则将输入栈的数据读取存入输出栈(导入）
	if len(this.stackOut) == 0 && len(this.stackIn) != 0 {
		for i := len(this.stackIn) - 1; i >= 0; i-- {
			// 输入栈入栈
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = make([]int, 0)
	}
	// 从stackOut栈顶取元素
	if len(this.stackOut) > 0 {
		x = this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
	}
	return x
}

// 返回队列开头的元素
func (this *MyQueue) Peek() int {
	x := 0
	// 如果输出栈为空, 则将输入栈的数据读取存入输出栈(导入）
	if len(this.stackOut) == 0 && len(this.stackIn) != 0 {
		for i := len(this.stackIn) - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = make([]int, 0)
	}
	// 从stackOut栈顶取元素
	if len(this.stackOut) > 0 {
		x = this.stackOut[len(this.stackOut)-1]
	}
	return x
}

// 如果队列为空，返回 `true` ；否则，返回 `false`
func (this *MyQueue) Empty() bool {
	if len(this.stackIn) != 0 || len(this.stackOut) != 0 {
		return false
	}
	return true
}

func main() {
	println("UseCase 1......")
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Printf("Peek: %v \n", queue.Peek())
	fmt.Printf("Pop: %v \n", queue.Pop())
	fmt.Printf("Empty: %v \n", queue.Empty())
}
