package main

import "fmt"

// 单调队列
type MonotonicQueue struct {
	queue []int //使用数组实现队列功能
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{queue: make([]int, 0)}
}

//(滑动窗口中移除元素的数值)
//如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
func (m *MonotonicQueue) Pop(value int) {
	if len(m.queue) > 0 && value == m.queue[0] {
		m.queue = m.queue[1:]
	}
}

//(滑动窗口添加元素的数值)
//如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
func (m *MonotonicQueue) Push(value int) {
	for len(m.queue) > 0 && value > m.queue[len(m.queue)-1] {
		m.queue = m.queue[:len(m.queue)-1]
	}
	// 添加到队尾
	m.queue = append(m.queue, value)
}

// 获取队列队首元素
func (m *MonotonicQueue) Front() int {
	if len(m.queue) > 0 {
		return m.queue[0]
	}
	return 0
}

// 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	maxArr := []int{}
	monotonicQueue := NewMonotonicQueue()
	// 先将第一个窗口元素全部塞入队列
	for i := 0; i < k; i++ {
		monotonicQueue.Push(nums[i])
	}
	maxArr = append(maxArr, monotonicQueue.Front())

	// 移动滑动窗口, 每次移动一位
	// 每次先优先级队列Pop出移除窗口的元素，然后在Push新进窗口的那个元素，最后获取此时窗口队首元素，即滑动窗口的最大值
	for i := k; i < numsLen; i++ {
		monotonicQueue.Pop(nums[i-k])
		monotonicQueue.Push(nums[i])
		maxArr = append(maxArr, monotonicQueue.Front())
	}
	return maxArr
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	println("UseCase 2......")
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
