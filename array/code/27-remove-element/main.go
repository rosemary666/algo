package main

import "fmt"

// 暴力破解法
func removeElement1(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val { // 数组值等于目标值，目标值后的元素整体就向前移动
			if i == size-1 {
				nums = nums[0:i]
			} else {
				nums = append(nums[0:i], nums[i+1:]...)
			}
			size -= 1 // 数组整体大小减1
			i -= 1    // 数组i也需要减1
		}
	}
	return len(nums)
}

func removeElement(nums []int, val int) int {
	var dest_index = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[dest_index] = nums[i]
			dest_index += 1
		}
	}
	return dest_index
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
