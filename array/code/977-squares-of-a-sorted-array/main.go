package main

import "fmt"

// 双指针法
func sortedSquares(nums []int) []int {
	size := len(nums)
	left_index := 0
	right_index := size - 1
	k := size - 1
	dest_arr := make([]int, size)

	for left_index <= right_index {
		left_sqrt := nums[left_index] * nums[left_index]
		right_sqrt := nums[right_index] * nums[right_index]
		if left_sqrt < right_sqrt {
			dest_arr[k] = right_sqrt
			right_index -= 1
		} else {
			dest_arr[k] = left_sqrt
			left_index += 1
		}
		k -= 1
	}
	return dest_arr
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
