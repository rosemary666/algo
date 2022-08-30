package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums1, 9))
	fmt.Println(search(nums2, 2))
}
