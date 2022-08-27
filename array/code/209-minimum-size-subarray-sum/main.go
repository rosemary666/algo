package main

import "fmt"

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	var (
		size               = len(nums)
		window_sum         = 0        //窗口值的综合
		min_window_length  = size + 1 //最小的窗口长度, 初始化可以初始化为数组长度+1
		window_start_index = 0        // 窗口开始索引
		window_end_index   = 0        // 窗口结束索引
		found              = false
	)
	for window_end_index = 0; window_end_index < size; window_end_index++ {
		window_sum += nums[window_end_index]
		// 如果滑动窗口总和超过目标值, 那么就不停的动态调整窗口开始位置
		for window_sum >= target {
			found = true
			window_length := window_end_index - window_start_index + 1
			if window_length < min_window_length {
				min_window_length = window_length
			}
			// 滑动窗口向前移动一位
			window_sum -= nums[window_start_index]
			window_start_index += 1
		}
	}
	if found {
		return min_window_length
	}
	return 0
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
