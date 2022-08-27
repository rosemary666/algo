package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums) // 数组先排序
	// 第一层循环
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		// n1去重
		if i > 0 && nums[i-1] == n1 {
			continue
		}
		// 第二层循环
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			// n2去重
			if j > i+1 && nums[j-1] == n2 {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			// 移动left和right指针直到相遇
			for left < right {
				n3 := nums[left]
				n4 := nums[right]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					// 像数值小的方向移动, 即rigth指针向左移动
					right--
				} else if sum < target {
					// 向数值大的方向，即向右移动
					left++
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					// left去重
					for left < right && nums[left] == n3 {
						left++
					}
					// right去重
					for left < right && nums[right] == n4 {
						right--
					}
				}
			}
		}
	}
	return res
}

func main() {
	println("UseCase 1......")
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Printf("%v\n", fourSum(nums, 0))

	println("UseCase 2......")
	nums = []int{2, 2, 2, 2, 2}
	fmt.Printf("%v\n", fourSum(nums, 8))
}
